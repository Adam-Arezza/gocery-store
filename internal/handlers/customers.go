package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
    "strconv"
    "fmt"
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

func ListCustomers(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var customers []Customer
    customersQuery := `SELECT * from customers;`
    rows, err := db.Query(customersQuery)

    if err != nil{
        log.Printf("Error fetching customers: %s\n", err.Error())
        http.Error(writer, "Server Error", http.StatusInternalServerError)
        return
    }

    defer rows.Close()

    for rows.Next(){
        var customer Customer        
        if err := rows.Scan(&customer.Id, &customer.Name, &customer.Email); err != nil{
            log.Printf("Error getting customers: %s\n", err.Error())
            http.Error(writer, "Server Error", http.StatusInternalServerError)
            return
        }
        customers = append(customers,customer)
    }
    writer.Header().Add("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(customers)
}

func GetCustomer(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var customer Customer
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
        http.Error(writer, "Server Error", http.StatusInternalServerError)
        return
    }
}

