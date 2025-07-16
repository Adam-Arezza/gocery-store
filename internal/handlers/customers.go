package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
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

    customer, err := services.CreateCustomer(db, newCustomer)

    http.Error(writer, "User already exists", http.StatusConflict)
}

func GetCustomers(writer http.ResponseWriter, r *http.Request, db *sql.DB){
    var customers []models.Customer

    //check for email query param
    email := r.URL.Query().Get("email")
    if email == ""{
        customersQuery := `SELECT * from customers;`
        rows, err := db.Query(customersQuery)

        if err != nil{
            log.Printf("Error fetching customers: %s\n", err.Error())
            http.Error(writer, "Server Error", http.StatusInternalServerError)
            return
        }

        defer rows.Close()

        for rows.Next(){
            var customer models.Customer        
            if err := rows.Scan(&customer.Id, &customer.Name, &customer.Email); err != nil{
                log.Printf("Error getting customers: %s\n", err.Error())
                http.Error(writer, "Server Error", http.StatusInternalServerError)
                return
            }
            customers = append(customers,customer)
        }
        writer.Header().Add("Content-Type", "application/json")
        json.NewEncoder(writer).Encode(customers)
        return
    }else{

        emailIsValid := validateEmail(email)

        if !emailIsValid{
            http.Error(writer, "Invalid request parameter: email", http.StatusBadRequest)
            return
        }

        var customer models.Customer
        customer, err := getCustomerByEmail(email, db)
        if err != nil && err == sql.ErrNoRows{
            log.Printf("No user found with email\n")
            noUserMsg := fmt.Sprintf("No user found with email: %s", email)
            http.Error(writer, noUserMsg, http.StatusNotFound)
            return
        }

        writer.Header().Add("Content-Type", "application/json")
        err = json.NewEncoder(writer).Encode(customer)

        if err != nil {
            log.Printf("Error in response: %s", err.Error())
            http.Error(writer, "Server Error", http.StatusInternalServerError)
            return
        }
        return
    }
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
        http.Error(writer, "Server Error", http.StatusInternalServerError)
        return
    }
}

//func getCustomerByEmail(email string, db *sql.DB)(models.Customer, error){
//    var customer models.Customer
//    customerQuery := `SELECT * FROM customers WHERE email = ?;`
//    err := db.QueryRow(customerQuery, email).Scan(&customer.Id,&customer.Name,&customer.Email)
//
//    if err != nil{
//        log.Printf("%s\n", err.Error())
//        return customer,err
//    }else{
//        return customer, nil
//    }
//}


//func validateEmail(email string)bool{
//    	// Define a regex pattern for email validation
//	// This is a simplified regex for basic email validation
//	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
//	
//	// Compile the regex
//	emailRegex := regexp.MustCompile(emailRegexPattern)
//	
//	// Return whether the email matches the pattern
//	return emailRegex.MatchString(email)
//}
//
//func checkIsExistingCustomer(customer models.Customer, db *sql.DB)(bool,error){
//    var existingCustomer models.Customer
//    checkCustomerQuery := `SELECT * FROM customers WHERE email = ?;`
//    err := db.QueryRow(checkCustomerQuery, customer.Email).Scan(&existingCustomer.Id,
//                                                                &existingCustomer.Name,
//                                                                &existingCustomer.Email)
//    if err != nil && err == sql.ErrNoRows{
//        return false,nil
//    }
//
//    if err != nil{
//        log.Printf("error checking existing customer: %s", err.Error())
//        return false, err
//    }
//
//    return true, nil
//}
