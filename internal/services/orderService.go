package services

import (
	"database/sql"
	"fmt"
	"time"
	"github.com/Adam-Arezza/gocery-store/internal/models"
    "log"
)

func CreateNewOrder(db *sql.DB, customer models.Customer) (sql.Result, error){
    var orderStatus models.OrderStatus
    var result sql.Result
    err := db.QueryRow(`SELECT * FROM order_statuses WHERE status='created';`).Scan(&orderStatus.Id, &orderStatus.Status)
    if err != nil && err == sql.ErrNoRows{
        return nil, fmt.Errorf(err.Error())
    }

    newOrder := models.Order{
        CustomerId: customer.Id,
        Date: time.Now().String(),
        StatusId: orderStatus.Id,
    }

    orderQuery := `INSERT INTO orders (customer_id, date, status_id) VALUES(?,?,?);`
    result, err = db.Exec(orderQuery, newOrder.CustomerId, newOrder.Date, orderStatus.Id)
    return result, nil
}

func GetOrders(db *sql.DB)([]models.Order, error){
    var orders []models.Order
    ordersQuery := `SELECT * from orders;`
    rows, err := db.Query(ordersQuery)

    if err != nil{
        log.Printf("Error fetching orders: %s\n", err.Error())
        return nil, fmt.Errorf(err.Error())
    }

    defer rows.Close()

    for rows.Next(){
        var order models.Order
        if err := rows.Scan(&order.Id, &order.CustomerId, &order.Date, &order.StatusId); err != nil{
            log.Printf("Error getting categories: %s\n", err.Error())
            return nil, fmt.Errorf(err.Error())
        }
        orders  = append(orders,order)
    }
    return orders, nil
}

func CancelOrder(db *sql.DB, orderId int)(bool){
    var order models.Order
    orderQuery := `SELECT * FROM orders WHERE id=?;`
    err := db.QueryRow(orderQuery, orderId).Scan(&order.Id,
                                                 &order.CustomerId,
                                                 &order.Date,
                                                 &order.StatusId)
    if err != nil {
        if err == sql.ErrNoRows{
            log.Printf("No order found for order ID: %v\n", orderId)
            return false
        }else{
            log.Printf(err.Error())
            return false
        }
    }else{
        status_id := models.StatusCancelled
        cancelOrderQuery := `UPDATE orders SET status_id=? WHERE id=?;`
        _, err := db.Exec(cancelOrderQuery,status_id,orderId)
        if err != nil {
            log.Printf(err.Error())
            return false
        }
        return true
    }
}

func AddItemsToOrder(orderItems []models.OrderItem, db *sql.DB){
    for _,v := range orderItems{
        if checkItemStock(v, db) <= v.Quantity{
            //there is enough stock to fulfill the addOrderItemsToOrder
            orderItemQuery := `INSERT INTO order_item (order_id, item_id, quantity) VALUES (?,?,?);`
            result, err := db.Exec(orderItemQuery, v.OrderId, v.ItemId, v.Quantity)
            if err != nil{
                log.Printf("Could not add order item: %v, %s",v,err.Error())
            }
            if id, err := result.LastInsertId(); err != nil{
                log.Printf("Error adding order item: %s", err.Error())
            }else{
                log.Printf("Added new order item with id: %v",id)
            }
        }else{
            //there is not enough stock to fulfill the order
            log.Printf("Not enough stock to fulfill the order.")
        }
    }
}

func GetOrderItems(orderItemReq models.OrderItemRequest, db *sql.DB)([] models.OrderRequestItem, error){
    var orderItems [] models.OrderRequestItem
    query := `SELECT
              g.name, 
              oi.quantity
              FROM order_item AS oi
              JOIN orders AS o ON oi.order_id = o.id
              JOIN customers AS c ON o.customer_id = c.id
              JOIN grocery_items AS g ON oi.item_id = g.id
              WHERE o.customer_id = ? AND oi.order_id = ?;`
    rows, err := db.Query(query, orderItemReq.CustomerId, orderItemReq.OrderId)
    
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    fmt.Println("Got rows...")
    for rows.Next(){
        var orderItem models.OrderRequestItem
        err := rows.Scan(&orderItem.Item, &orderItem.Quantity)
        if err != nil{
            return nil, err
        }

        orderItems = append(orderItems, orderItem)
    }

    return orderItems, nil
}

func checkItemStock(orderItem models.OrderItem, db *sql.DB)int{
    var item models.GroceryItem
    query := `SELECT stock FROM grocery_items WHERE item_id = ?;`
    err := db.QueryRow(query, orderItem.ItemId).Scan(&item)
    if err != nil{
        if err == sql.ErrNoRows{
            log.Printf("No grocery item found: %v, %s\n", item.Id, item.Name)
            return 0
        }else{
            log.Printf(err.Error())
            return 0
        }
    }
    return item.Stock
}

func IsCustomersOrder(customerId int, orderId int, db *sql.DB)(bool){
    query := `SELECT * FROM orders WHERE customer_id = ? AND id = ?;`
    _,err := db.Query(query, customerId, orderId)
    if err != nil && err == sql.ErrNoRows{
        log.Printf("No orders found for customer")
        return false
    }
    if err != nil{
        log.Printf("%s", err.Error())
        return false
    }
    return true
}
