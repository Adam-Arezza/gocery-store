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
    //TODO trailing slashes don't work with URL query strings
    router.HandleFunc("GET /grocery_items", func(w http.ResponseWriter, r *http.Request){
        handlers.GetGroceryItemsHandler(w,r,db)
    })

    router.HandleFunc("GET /grocery_items/{id}", func(w http.ResponseWriter, r *http.Request){
        handlers.GetGroceryItemByIdHandler(w,r,db)
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

    //orders

    router.HandleFunc("GET /orders", func(w http.ResponseWriter, r *http.Request){
        handlers.GetOrders(w,r,db)
    })

    router.HandleFunc("POST /orders", func(w http.ResponseWriter, r *http.Request){
        handlers.CreateOrder(w,r,db)
    })

    router.HandleFunc("PUT /orders/{id}", func(w http.ResponseWriter, r *http.Request){
        handlers.UpdateOrder(w,r,db)
    })

    router.HandleFunc("PUT /orders/{id}/cancel", func(w http.ResponseWriter, r *http.Request){
        handlers.CancelOrder(w,r,db)
    })

    router.HandleFunc("GET /order_items", func(w http.ResponseWriter, r *http.Request){
        handlers.GetOrderItems(w,r,db)
    })
    

    return router
}
