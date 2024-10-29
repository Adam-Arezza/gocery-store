package main

import (
	"fmt"
	"net/http"
	"github.com/Adam-Arezza/gocery-store/internal/database"
	"github.com/Adam-Arezza/gocery-store/internal/routes"
)


func main(){
    db, err := database.DbInit()

    if err != nil{
        fmt.Println("Error opening database")
        fmt.Println(err)
        return
    }

    defer db.Close()

    router := routes.SetRoutes(db)
    fmt.Println("Starting server on port 3000...")
    if err := http.ListenAndServe(":3000", router); err != nil{
        fmt.Printf("Error starting server: \n%s", err)
    }

    return
}

