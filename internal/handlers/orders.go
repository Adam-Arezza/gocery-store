package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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


//update order
//cancel order

