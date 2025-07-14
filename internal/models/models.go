package models

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
    ItemId int `json:"item_id"`
    Quantity int `json:"quantity"`
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
    Status string `json:"status"`
}

type UpdateOrderRequest struct{
    CustomerId int `json:"customer_id"`
    ItemList []OrderRequestItem
    OrderId int `json:"order_id"`
}

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
