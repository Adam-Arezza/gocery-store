package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
    "github.com/Adam-Arezza/gocery-store/internal/models"
    "github.com/Adam-Arezza/gocery-store/internal/services"
)

//create order
func CreateOrder(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var customer models.Customer 
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    err := decoder.Decode(&customer)

    if err != nil {
        log.Printf("Error with request body: %s", err.Error())
        http.Error(writer, "Error processing request", http.StatusBadRequest)
        return
    }   

    isExistingCustomer, err := services.CheckIsExistingCustomer(customer, db)
    if err != nil{
        http.Error(writer, "Server Error", http.StatusInternalServerError)
        return
    }

    if !isExistingCustomer{
        log.Printf("Customer doesn't exist, order creation aborted")
        http.Error(writer, "Customer doesn't exist, order creation aborted. Please create customer.", http.StatusBadRequest)
        return
    }

    result,err := services.CreateNewOrder(db, customer)     
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
        log.Printf("New order created with id: %d", int(newOrderId))
        json.NewEncoder(writer).Encode(newOrderId)
        return
    }
}

func UpdateOrder(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    //get the order items from the request body
    var orderItems []models.OrderItem
    var orderRequest models.UpdateOrderRequest
    orderId, err := strconv.Atoi(r.PathValue("id"))
    if err != nil {
        log.Printf("Error with order id in request")
    }
    log.Printf("The order id is: %d", orderId)
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    err = decoder.Decode(&orderRequest)
    if err != nil{
        log.Printf("Error parsing request: %s", err.Error())
        http.Error(writer, "Error parsing request", http.StatusBadRequest)
        return
    }
    
    //validate the order belongs to the customer making the change
    validOrder := services.IsCustomersOrder(orderRequest.CustomerId, orderRequest.OrderId, db)
    if !validOrder{
        log.Printf("Invalid order request")
        http.Error(writer, "The order was invalid", http.StatusBadRequest)
        return
    }

    for item := range orderRequest.ItemList {
        newOrderItem := models.OrderItem{
            ItemId: orderRequest.ItemList[item].ItemId, 
            OrderId: orderRequest.OrderId, 
            Quantity: orderRequest.ItemList[item].Quantity,
        }
        orderItems = append(orderItems, newOrderItem)
    }
    services.AddItemsToOrder(orderItems, db)
}

//get orders by customer
func GetOrders(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    orders, err := services.GetOrders(db)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
    log.Print(orders)
    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(orders)
}

func GetOrderItems(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var orderItemRequest models.OrderItemRequest
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    err := decoder.Decode(&orderItemRequest)

    if err != nil{
        http.Error(writer, err.Error(), http.StatusBadRequest)
        return
    }

    orderItems, err := services.GetOrderItems(orderItemRequest, db)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusNotFound)
        return
    }

    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(&orderItems)
}

//cancel order 
func CancelOrder(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    orderId, err := strconv.Atoi(r.PathValue("id"))
    if err != nil {
        log.Printf("Error with order id in request")
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
    orderCancelled := services.CancelOrder(db, orderId)
    if orderCancelled {
        writer.WriteHeader(http.StatusNoContent)
        return 
    }else{
        writer.WriteHeader(http.StatusNotFound)
        return
    }
}

