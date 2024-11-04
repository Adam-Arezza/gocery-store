package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateCustomer(writer http.ResponseWriter, r *http.Request, db *sql.DB){
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

    checkCustomerQuery := `SELECT * from customers WHERE email = ?;`
    err = db.QueryRow(checkCustomerQuery, newCustomer.Email).Scan(&customer.Id,&customer.Name,&customer.Email)

    //no customer was found, try to add new one
    if err == sql.ErrNoRows{
        var result sql.Result
        log.Println("No rows in result") 
        log.Printf("Adding new customer: %s, %s\n", newCustomer.Name, newCustomer.Email)
        newCustomerQuery := `INSERT INTO customers (name, email) VALUES(?,?);`
        result, err = db.Exec(newCustomerQuery, newCustomer.Name, newCustomer.Email)

        if err != nil{
            log.Println("error executing add new user query")
            return
        }

        //check result of query
        if rows, err := result.RowsAffected(); err != nil || rows < 1{
            log.Println("error with query result")
            log.Printf("%s\n", err)
            return
        }else{
            rowId, _ := result.LastInsertId()
            newCustomer.Id = int(rowId)
            json.NewEncoder(writer).Encode(newCustomer)
            return
        }
    }

    //there is an error and no customer found
    if err != nil{
        log.Printf("Error checking for customer: %s", err.Error())
        http.Error(writer, "Server error", http.StatusInternalServerError)
        return
    }

    //customer was found already
    http.Error(writer, "User already exists", http.StatusConflict)
}
