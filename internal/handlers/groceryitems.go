package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "database/sql"
)

type GroceryItem struct {
    Id int `json:"id"`
    Name string `json:"item_name"`
    UnitPrice float64 `json:"unit_price"`
    Stock int `json:"stock"`
    CategoryId int `json:"category_id"`
}

func ListGroceryItems(writer http.ResponseWriter, r *http.Request, db *sql.DB){
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

func GetGroceryItem(writer http.ResponseWriter, r *http.Request, db *sql.DB){
}
