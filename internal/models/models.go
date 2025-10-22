package models

type Status int

const (
    _ Status = iota
    StatusCompleted 
    StatusCreated     
    StatusCancelled 
)

var statusString = map[Status]string{
    StatusCreated: "created",
    StatusCompleted: "completed",
    StatusCancelled: "cancelled",
}


type GroceryItem struct {
    Id int `json:"id"`
    Name string `json:"item_name"`
    UnitPrice float64 `json:"unit_price"`
    Stock int `json:"stock"`
    CategoryId int `json:"category_id"`
}

type Category struct {
    Id int `json:"id"`
    Name string `json:"category"`
}

type OrderRequestItem struct {
    Item string `json:"item"`
    Quantity int `json:"quantity"`
}

type OrderItemRequest struct {
    CustomerId int `json:"customer_id"`
    OrderId int `json:"order_id"`
}

type Order struct {
    Id int `json:"id"`
    CustomerId int `json:"customer_id"`
    Date string `json:"date"`
    StatusId int `json:"order_status"`
}

type OrderItem struct{
    Id int `json:"id"`
    OrderId int `json:"order_id"`
    ItemId int `json:"item_id"`
    Quantity int `json:"quantity"`
}

type OrderStatus struct{
    Id int `json:"id"`
    Status Status `json:"status"`
}

type UpdateOrderItems struct {
    ItemId int `json:"item_id"`
    Quantity int `json:"quantity"`
}

type UpdateOrderRequest struct{
    CustomerId int `json:"customer_id"`
    ItemList []UpdateOrderItems `json:"order_request_item"`
    OrderId int `json:"order_id"`
}

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
