package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Category struct {
    Id int `json:"id"`
    Name string `json:"category"`
}

func ListCategories(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var categories []Category
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


func GetCategory(writer http.ResponseWriter, r *http.Request, db *sql.DB)(error){
    var category Category
    categoryId, err := strconv.Atoi(r.PathValue("id"))
    fmt.Printf("Category id: %d\n", categoryId)

    if err != nil{
        return fmt.Errorf("Error getting category: %s", err)
    }

    categoryQuery := `SELECT * FROM categories WHERE id = ?;`
    err = db.QueryRow(categoryQuery, categoryId).Scan(&category.Id,&category.Name)

    if err != nil{
        fmt.Printf("error in query: %s\n",err)
        return fmt.Errorf("Error getting category: %s", err)
    }

    writer.Header().Add("Content-Type", "application/json")
    fmt.Printf("The category is: %s", category.Name)
    err = json.NewEncoder(writer).Encode(category)

    if err != nil {
        return fmt.Errorf("Error writing json data: %s", err)
    }

    return nil
}


