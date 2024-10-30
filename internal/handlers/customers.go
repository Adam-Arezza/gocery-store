package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
    "fmt"
)

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//create user
//update user
//delete user
//list users

func CreateCustomer(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    contentType := r.Header.Get("Content-Type")
    fmt.Printf("content type: %s", contentType)
    var newCustomer Customer
    var customer Customer
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    err := decoder.Decode(&newCustomer)
    if err != nil {
        log.Printf("Error with request body: %s", err.Error())
        http.Error(writer, "Error processing request", http.StatusBadRequest)
        return
    }

    //database
    //check if customer already exists (don't create a duplicate)
    //add the new customer to the db
    //return successful http response with the customer
    //or an http error
    checkCustomerQuery := `SELECT * from customers WHERE email = ?;`
    err = db.QueryRow(checkCustomerQuery, customer.Email).Scan(&customer.Id,&customer.Name,&customer.Email)

    //no customer was found, add new one
    if err == sql.ErrNoRows{
        log.Println("No rows in result") 
        newCustomerQuery := `INSERT INTO customers (name, email) VALUES(?,?);`
        err = db.QueryRow(newCustomerQuery, customer.Name, customer.Email).Scan(&newCustomer.Id,&newCustomer.Name, &newCustomer.Email)

    }

    //there is an error and no customer found
    if err != nil && err != sql.ErrNoRows{
        log.Printf("Error checking for customer: %s", err.Error())
        http.Error(writer, "Server error", http.StatusInternalServerError)
        return
    }

    //customer was found already
    http.Error(writer, "User already exists", http.StatusConflict)
}
