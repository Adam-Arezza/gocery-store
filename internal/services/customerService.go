package services

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"github.com/Adam-Arezza/gocery-store/internal/models"
)


func CreateCustomer(db *sql.DB, newCustomer models.Customer)(error){
    isExistingCustomer, err := CheckIsExistingCustomer(newCustomer, db)
    if err != nil {
        return fmt.Errorf("Error retreiving customer: %s", err.Error())
    }

    if !isExistingCustomer{
        var result sql.Result
        log.Println("No rows in result") 
        log.Printf("Adding new customer: %s, %s\n", newCustomer.Name, newCustomer.Email)
        newCustomerQuery := `INSERT INTO customers (name, email) VALUES(?,?);`
        result, err = db.Exec(newCustomerQuery, newCustomer.Name, newCustomer.Email)

        if err != nil{
            log.Println("error executing add new user query")
            return fmt.Errorf("Couldn't create customer: %s", err.Error())
        }

        //check result of query
        if rows, err := result.RowsAffected(); err != nil || rows < 1{
            log.Println("error with query result")
            log.Printf("%s\n", err)
            return err
        }
    }
    return nil
}

func GetCustomer(db *sql.DB, email string) ([]models.Customer, error){
    var customers []models.Customer
    customersQuery := `SELECT * from customers WHERE 1=1`
    isValidEmail := validateEmail(email)

    if email != "" && !isValidEmail{
        err := fmt.Errorf("Invalid Email Error")
        return nil, err
    }

    if email != "" && isValidEmail{
        customersQuery = fmt.Sprintf("%s AND email = '%s'", customersQuery, email)
    }

    rows, err := db.Query(customersQuery)
    if err != nil{
        log.Printf("Error fetching customers: %s\n", err.Error())
        return nil, err
    }

    defer rows.Close()

    for rows.Next(){
        var customer models.Customer        
        if err := rows.Scan(&customer.Id, &customer.Name, &customer.Email); err != nil{
            log.Printf("Error getting customers: %s\n", err.Error())
            return nil, err
        }
        customers = append(customers,customer)
    }
    return customers, nil
}

func GetCustomerById(db *sql.DB, id int) ([]models.Customer, error){
    var customer [] models.Customer
    customerQuery := fmt.Sprintf(`SELECT * from customers WHERE id = %d`, id)
    rows, err := db.Query(customerQuery)
    defer rows.Close()
    if err != nil{
        return nil, err
    }

    for rows.Next(){
        var resultCustomer models.Customer        
        if err := rows.Scan(&resultCustomer.Id, &resultCustomer.Name, &resultCustomer.Email); err != nil{
            log.Printf("Error getting customers: %s\n", err.Error())
            return nil, err
        }
        customer = append(customer,resultCustomer)
    }
    return customer, nil
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

func CheckIsExistingCustomer(customer models.Customer, db *sql.DB)(bool,error){
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
