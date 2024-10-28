package routes

import (
	"database/sql"
	"net/http"
	"github.com/Adam-Arezza/gocery-store/internal/handlers"
)


func SetRoutes(db *sql.DB) *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc("GET /categories", func(w http.ResponseWriter, r *http.Request){
        handlers.ListCategories(w,r,db)
    })

    router.HandleFunc("GET /categories/{id}", func(w http.ResponseWriter, r *http.Request){
        handlers.GetCategory(w,r,db)
    })

    router.HandleFunc("GET /grocery_items", func(w http.ResponseWriter, r *http.Request){
        handlers.ListGroceryItems(w,r,db)
    })

    router.HandleFunc("GET /grocery_items/{id}", func(w http.ResponseWriter, r *http.Request){
        handlers.GetGroceryItem(w,r,db)
    })

    return router
}
