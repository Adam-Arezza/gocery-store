package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
    "fmt"
)

type GroceryItem struct {
    Id int `json:"id"`
    Name string `json:"item_name"`
    UnitPrice float64 `json:"unit_price"`
    Stock int `json:"stock"`
    CategoryId int `json:"category_id"`
}

func GetGroceryItems(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var groceries []GroceryItem
    groceryName := r.URL.Query().Get("name")
    if groceryName == ""{
        groceryQuery := `SELECT * from grocery_items;`
        rows, err := db.Query(groceryQuery)

        if err != nil{
            log.Printf("Error getting grocery items: %s", err.Error())
            http.Error(writer, "Server Error", http.StatusInternalServerError)
            return
        }

        defer rows.Close()

        for rows.Next(){
            var groceryItem GroceryItem

            err := rows.Scan(&groceryItem.Id, &groceryItem.Name, &groceryItem.UnitPrice, &groceryItem.Stock, &groceryItem.CategoryId);
            if err != nil{
                log.Printf("Error reading database: %s", err.Error())
                http.Error(writer, "Server Error", http.StatusInternalServerError)
                return
            }
            groceries = append(groceries, groceryItem)
        }
        writer.Header().Add("Content-Type", "application/json")
        json.NewEncoder(writer).Encode(groceries)
        return
    }else{
                //add some kind of validation for the grocery names
        var groceryItem GroceryItem
        var err error
        groceryItem, err = getGroceryItemByName(groceryName, db)
        if err != nil && err == sql.ErrNoRows{
            log.Printf("No grocery item found with name\n")
            noItemMsg := fmt.Sprintf("No grocery item with name: %s", groceryName)
            http.Error(writer, noItemMsg, http.StatusNotFound)
            return
        }

        writer.Header().Add("Content-Type", "application/json")
        err = json.NewEncoder(writer).Encode(groceryItem)
        if err != nil{
            log.Printf("Error in response: %s", err.Error())
            http.Error(writer, "Server error", http.StatusInternalServerError)
            return
        }
    }
}

func GetGroceryItemById(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var groceryItem GroceryItem
    groceryItemId, err := strconv.Atoi(r.PathValue("id"))
    if err != nil{
        log.Printf("Invalid URL path value 'id': %s", r.PathValue("id"))
        http.Error(writer, "Invvalid URL path value 'id'", http.StatusBadRequest)
        return
    }

    groceryQuery := `SELECT * from grocery_items WHERE id = ?;`

    err = db.QueryRow(groceryQuery, groceryItemId).Scan(&groceryItem.Id,&groceryItem.Name,&groceryItem.UnitPrice,&groceryItem.Stock,&groceryItem.CategoryId)

    if err != nil{
        log.Printf("%s\n", err.Error())
        http.Error(writer, "Couldn't find grocery item", http.StatusNotFound)
        return
    }
    writer.Header().Add("Content-Type", "application/json")
    err = json.NewEncoder(writer).Encode(groceryItem)

    if err != nil {
        log.Printf("Error in response: %s", err.Error())
        http.Error(writer, "Server Error", http.StatusInternalServerError)
        return
    }
}

func getGroceryItemByName(name string, db *sql.DB)(GroceryItem, error){
    var groceryItem GroceryItem
    groceryQuery := fmt.Sprintf(`SELECT * FROM grocery_items WHERE name='%s'`, name)
    err := db.QueryRow(groceryQuery, name).Scan(&groceryItem.Id, &groceryItem.Name, &groceryItem.UnitPrice, &groceryItem.Stock, &groceryItem.CategoryId)
    if err != nil{
        log.Println(groceryQuery)
        log.Printf("Error getting grocery item: %s",err.Error())
        return groceryItem, err
    }else{
        return groceryItem, nil
    }
}
