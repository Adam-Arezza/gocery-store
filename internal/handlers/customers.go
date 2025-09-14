package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/Adam-Arezza/gocery-store/internal/models"
	"github.com/Adam-Arezza/gocery-store/internal/services"
)

func CreateCustomer(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var newCustomer models.Customer
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    err := decoder.Decode(&newCustomer)
    if err != nil {
        log.Printf("Error with request body: %s", err.Error())
        http.Error(writer, "Error processing request", http.StatusBadRequest)
        return
    }

    err = services.CreateCustomer(db, newCustomer)
    if err != nil {
        log.Printf("Error with request body: %s", err.Error())
        http.Error(writer, "Error processing request", http.StatusBadRequest)
        return
    }
    json.NewEncoder(writer).Encode(newCustomer)
    return
}

func GetCustomers(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    email := r.URL.Query().Get("email")
    customers,err := services.GetCustomer(db, email)
    if(err != nil){
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(customers)
    return
}

func GetCustomerById(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var customer models.Customer
    customerId, err := strconv.Atoi(r.PathValue("id"))

    if err != nil{
        log.Printf("Ivalid URL path value 'id': %s\n Error: %s", r.PathValue("id"), err.Error())
        http.Error(writer, fmt.Sprintf("Ivalid URL path value 'id': %s", r.PathValue("id")), http.StatusBadRequest)
        return
    }

    customerQuery := `SELECT * FROM customers WHERE id = ?;`
    err = db.QueryRow(customerQuery, customerId).Scan(&customer.Id,&customer.Name,&customer.Email)

    if err != nil{
        log.Printf("%s\n", err.Error())
        http.Error(writer, "Couldn't find Customer", http.StatusNotFound)
        return
    }

    writer.Header().Add("Content-Type", "application/json")
    err = json.NewEncoder(writer).Encode(customer)

    if err != nil {
        log.Printf("Error in response: %s", err.Error())
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
}
