package services

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"github.com/Adam-Arezza/gocery-store/internal/models"
)


func CreateCustomer(db *sql.DB, newCustomer models.Customer)([]models.Customer, error){
    isExistingCustomer, err := checkIsExistingCustomer(newCustomer, db)
    if err != nil {
        return nil, fmt.Errorf("Error retreiving customer: %s", err.Error())
    }

    if !isExistingCustomer{
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
}

func GetCustomer(){

}

func GetCustomerById(){

}

func getCustomerByEmail(email string, db *sql.DB)(models.Customer, error){
    var customer models.Customer
    customerQuery := `SELECT * FROM customers WHERE email = ?;`
    err := db.QueryRow(customerQuery, email).Scan(&customer.Id,&customer.Name,&customer.Email)

    if err != nil{
        log.Printf("%s\n", err.Error())
        return customer,err
    }else{
        return customer, nil
    }
}

func validateEmail(email string)bool{
    	// Define a regex pattern for email validation
	// This is a simplified regex for basic email validation
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	
	// Compile the regex
	emailRegex := regexp.MustCompile(emailRegexPattern)
	
	// Return whether the email matches the pattern
	return emailRegex.MatchString(email)
}

func checkIsExistingCustomer(customer models.Customer, db *sql.DB)(bool,error){
    var existingCustomer models.Customer
    checkCustomerQuery := `SELECT * FROM customers WHERE email = ?;`
    err := db.QueryRow(checkCustomerQuery, customer.Email).Scan(&existingCustomer.Id,
                                                                &existingCustomer.Name,
                                                                &existingCustomer.Email)
    if err != nil && err == sql.ErrNoRows{
        return false, nil
    }

    if err != nil{
        log.Printf("error checking existing customer: %s", err.Error())
        return false, err
    }

    return true, nil
}
