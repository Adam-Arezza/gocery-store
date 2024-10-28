package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
    "encoding/json"
)


type Category struct {
    Id int `json:"id"`
    Name string `json:"category"`
}

type GroceryItem struct {
    Id int `json:"id"`
    Name string `json:"item_name"`
    UnitPrice float64 `json:"unit_price"`
    Stock int `json:"stock"`
    CategoryId int `json:"category_id"`
}

type Order struct {
    Id int
    CustomerId int
    Date string
}

type Customer struct {
    Id int
    Name string
    Email string
}

type OrderItem struct{
    Id int
    OrderId int
    ItemId int
    Quantity int
}


func CategoriesHandler(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    
    var categories []Category
    fmt.Println(r.URL)
    categoriesQuery := `SELECT * from categories;`
    rows, err := db.Query(categoriesQuery)

//need to change how the errors are handled here, shouldn't pass them to the user
    if err != nil{
        fmt.Printf("Error getting categories: %s", err)
        fmt.Fprint(writer, fmt.Errorf("There was an error getting the categories..."))
    }

    defer rows.Close()

    for rows.Next(){
        var category Category
        if err := rows.Scan(&category.Id, &category.Name); err != nil{
            fmt.Printf("Error getting categories: %s", err)
            fmt.Fprint(writer, fmt.Errorf("There was an error getting the categories..."))

        }
        categories = append(categories,category)
    }
    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(categories)
}


func GroceryItemsHandler(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var groceries []GroceryItem
    groceryQuery := `SELECT * from grocery_items;`
    rows, err := db.Query(groceryQuery)

    if err != nil{
        fmt.Printf("Error getting grocery items: %s", err)
    }

    defer rows.Close()

    for rows.Next(){
        var groceryItem GroceryItem

        if err:= rows.Scan(&groceryItem.Id, 
                           &groceryItem.Name, 
                           &groceryItem.UnitPrice, 
                           &groceryItem.Stock, 
                           &groceryItem.CategoryId);
                           err != nil{
                               fmt.Printf("Error getting grocery items: %s", err)
                           }
                           groceries = append(groceries, groceryItem)
    }
    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(groceries)
}

