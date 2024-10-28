package routes

import (
	"database/sql"
	"net/http"
	"github.com/Adam-Arezza/gocery-store/internal/handlers"
)


func SetRoutes(db *sql.DB) *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc("GET /categories", func(w http.ResponseWriter, r *http.Request){
        handlers.CategoriesHandler(w,r,db)
    })

    router.HandleFunc("GET /grocery_items", func(w http.ResponseWriter, r *http.Request){
        handlers.GroceryItemsHandler(w,r,db)
    })

    return router
}
