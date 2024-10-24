package main

import (
	"fmt"

	"github.com/Adam-Arezza/gocery-store/internal/database"
)

func main(){
    db, err := database.DbInit()
    if err != nil{
        fmt.Println("Error opening database")
    }
    db.Close()
    return
}
