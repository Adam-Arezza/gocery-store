package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type OrderRequestItem struct {
    ItemId int
    Quantity int
}

type Order struct {
    Id int `json:"id"`
    CustomerId int `json:"customer_id"`
    Date string `json:"date"`
    Status string `json:"order_status"`
}

type OrderItem struct{
    Id int `json:"id"`
    OrderId int `json:"order_id"`
    ItemId int `json:"item_id"`
    Quantity int `json:"quantity"`
}

type OrderStatus struct{
    Id int `json:"id"`
    Status string `json:"status"`
}

type UpdateOrderRequest struct{
    CustomerId int
    ItemList []OrderRequestItem
    OrderId int
}

//create order
func CreateOrder(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var customer Customer
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    err := decoder.Decode(&customer)

    if err != nil {
        log.Printf("Error with request body: %s", err.Error())
        http.Error(writer, "Error processing request", http.StatusBadRequest)
        return
    }   

    isExistingCustomer, err := checkIsExistingCustomer(customer, db)
    if err != nil{
        http.Error(writer, "Server Error", http.StatusInternalServerError)
        return
    }

    if !isExistingCustomer{
        log.Printf("Customer doesn't exist, order creation aborted")
        http.Error(writer, "Customer doesn't exist, order creation aborted. Please create customer.", http.StatusBadRequest)
        return
    }

    var orderStatus OrderStatus
    var result sql.Result
    err = db.QueryRow(`SELECT * FROM order_statuses WHERE status='created';`).Scan(&orderStatus.Id, &orderStatus.Status)
    if err != nil && err == sql.ErrNoRows{
        log.Printf(err.Error())
        http.Error(writer, "Server Error", http.StatusInternalServerError)
        return
    }

    newOrder := Order{
        CustomerId: customer.Id,
        Date: time.Now().String(),
        Status: orderStatus.Status,
    }

    orderQuery := `INSERT INTO orders (customer_id, date, status_id) VALUES(?,?,?);`
    result, err = db.Exec(orderQuery, newOrder.CustomerId, newOrder.Date, orderStatus.Id)
    if err != nil{
        log.Printf(err.Error())
        http.Error(writer, "Could not create order", http.StatusInternalServerError)
        return
    }

    if newOrderId,err := result.LastInsertId();err != nil{
        log.Printf(err.Error())
        http.Error(writer, "Could not create new order", http.StatusInternalServerError)
        return
    }else{
        newOrder.Id = int(newOrderId)
        log.Printf("New order created with id: %d", newOrder.Id)
        json.NewEncoder(writer).Encode(newOrder)
        return
    }
}

func UpdateOrder(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    //get the order items from the request body
    var orderItems []OrderItem
    var orderRequest UpdateOrderRequest
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    err := decoder.Decode(&orderRequest)
    if err != nil{
        log.Printf("Error parsing request: %s", err.Error())
        http.Error(writer, "Error parsing request", http.StatusBadRequest)
        return
    }
    
    //validate the order belongs to the customer making the change
    validOrder := orderBelongsToCustomer(orderRequest.CustomerId, orderRequest.OrderId, db)
    if !validOrder{
        log.Printf("Invalid order request")
        http.Error(writer, "The order was invalid", http.StatusBadRequest)
        return
    }

    for item := range orderRequest.ItemList {
        newOrderItem := OrderItem{ItemId: orderRequest.ItemList[item].ItemId, 
                                  OrderId: orderRequest.OrderId, 
                                  Quantity: orderRequest.ItemList[item].Quantity}
        orderItems = append(orderItems, newOrderItem)
    }
    addOrderItemsToOrder(orderItems, db)
}

//get orders by customer

func GetOrders(writer http.ResponseWriter, r *http.Request, db *sql.DB){

}

//cancel order 
func CancelOrder(writer http.ResponseWriter, r *http.Request, db *sql.DB){

}

func orderBelongsToCustomer(customerId int, orderId int, db *sql.DB)(bool){
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

func addOrderItemsToOrder(orderItems []OrderItem, db *sql.DB){
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

func checkItemStock(orderItem OrderItem, db *sql.DB)int{
    var item GroceryItem
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

