package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
    "github.com/Adam-Arezza/gocery-store/internal/models"
)

func GetCategories(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var categories []models.Category
    categoriesQuery := `SELECT * from categories;`
    rows, err := db.Query(categoriesQuery)

    if err != nil{
        log.Printf("Error fetching categories: %s\n", err.Error())
        http.Error(writer, "Server Error", http.StatusInternalServerError)
        return
    }

    defer rows.Close()

    for rows.Next(){
        var category models.Category
        if err := rows.Scan(&category.Id, &category.Name); err != nil{
            log.Printf("Error getting categories: %s\n", err.Error())
            http.Error(writer, "Server Error", http.StatusInternalServerError)
            return
        }
        categories = append(categories,category)
    }
    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(categories)
}


func GetCategoryById(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var category models.Category
    categoryId, err := strconv.Atoi(r.PathValue("id"))

    if err != nil{
        log.Printf("Ivalid URL path value 'id': %s\n Error: %s", r.PathValue("id"), err.Error())
        http.Error(writer, fmt.Sprintf("Ivalid URL path value 'id': %s", r.PathValue("id")), http.StatusBadRequest)
        return
    }

    categoryQuery := `SELECT * FROM categories WHERE id = ?;`
    err = db.QueryRow(categoryQuery, categoryId).Scan(&category.Id,&category.Name)

    if err != nil{
        log.Printf("%s\n", err.Error())
        http.Error(writer, "Couldn't find Category", http.StatusNotFound)
        return
    }

    writer.Header().Add("Content-Type", "application/json")
    err = json.NewEncoder(writer).Encode(category)

    if err != nil {
        log.Printf("Error in response: %s", err.Error())
        http.Error(writer, "Server Error", http.StatusInternalServerError)
        return
    }
}


