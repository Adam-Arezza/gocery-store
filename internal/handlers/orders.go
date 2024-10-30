package handlers

//import (
//    "fmt"
//    "encoding/json"
//    "net/http"
//    "database/sql"
//)

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

//create order
//update order
//cancel order

