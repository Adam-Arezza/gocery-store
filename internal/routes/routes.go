package routes

import (
	"database/sql"
	"net/http"
	"github.com/Adam-Arezza/gocery-store/internal/handlers"
)


func SetRoutes(db *sql.DB) *http.ServeMux {
    router := http.NewServeMux()

    //categories
    router.HandleFunc("GET /categories", func(w http.ResponseWriter, r *http.Request){
        handlers.GetCategories(w,r,db)
    })

    router.HandleFunc("GET /categories/{id}", func(w http.ResponseWriter, r *http.Request){
        handlers.GetCategoryById(w,r,db)
    })

    //grocery_items
    router.HandleFunc("GET /grocery_items", func(w http.ResponseWriter, r *http.Request){
        handlers.GetGroceryItems(w,r,db)
    })

    router.HandleFunc("GET /grocery_items/{id}", func(w http.ResponseWriter, r *http.Request){
        handlers.GetGroceryItemById(w,r,db)
    })

    //customers
    router.HandleFunc("GET /customers", func(w http.ResponseWriter, r *http.Request){
        handlers.GetCustomers(w,r,db)
    })

    router.HandleFunc("POST /customers", func(w http.ResponseWriter, r *http.Request){
        handlers.CreateCustomer(w,r,db)
    })

    router.HandleFunc("GET /customers/{id}", func(w http.ResponseWriter, r *http.Request){
        handlers.GetCustomerById(w,r,db)
    })

    return router
}
