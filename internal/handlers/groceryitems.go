package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Adam-Arezza/gocery-store/internal/models"
	"github.com/Adam-Arezza/gocery-store/internal/services"
)

func GetGroceryItemsHandler(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    groceryName := r.URL.Query().Get("name")
    if groceryName != ""{
        groceries, err := services.GetGroceryItemByName(groceryName, db)
        if err != nil{
            fmt.Printf("No grocery item with name %s", groceryName)
            http.Error(writer, "No item found", http.StatusNotFound)
            return
        }
        json.NewEncoder(writer).Encode(groceries)
        return
    }else{
        writer.Header().Add("Content-Type", "application/json")
        groceries, err := services.GetGroceryItems(db)
        if err != nil {
            http.Error(writer, "Server Error", http.StatusInternalServerError)
            return 
        }
        json.NewEncoder(writer).Encode(groceries)
        return
    }
}

func GetGroceryItemByIdHandler(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    groceryItemId, err := strconv.Atoi(r.PathValue("id"))
    if err != nil{
    }
    groceries, err := services.GetGroceryItemById(groceryItemId, db)
    json.NewEncoder(writer).Encode(groceries)
    return
}

func UpdateGroceryItemHandler(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var groceryUpdateReq []models.UpdateGroceryStockRequest
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    err := decoder.Decode(&groceryUpdateReq)

    if err != nil {
        http.Error(writer, err.Error(), http.StatusBadRequest)
        return
    }

    err = services.UpdateGroceryItem(db, groceryUpdateReq)

    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
}
