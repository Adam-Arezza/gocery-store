package database

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	_ "github.com/mattn/go-sqlite3"
)


func DbInit() (*sql.DB, error){
    var databasePath string
    dir, err := os.UserHomeDir()
    if err != nil{
        return nil, fmt.Errorf("There was an error: %s", err)
    }

    if runtime.GOOS == "windows"{
        databasePath = dir + "\\grocery_store\\data"
    }else{
        databasePath = dir + "/grocery_store/data"
    }

    if _, err := os.Stat(databasePath); err == nil{
        fmt.Printf("Database exists at: %s\n", databasePath)
        db, err := sql.Open("sqlite3", databasePath)
        if err != nil{
            fmt.Println("Error creating database:")
            fmt.Println(err)
            return nil, fmt.Errorf("There was an error creating the database: %s", err)
        }
        return db, nil
    } else{
        fmt.Println("Database doesn't exist, creating now...")
        if err := os.MkdirAll(databasePath,0777); err != nil{
            return nil, fmt.Errorf("Error creating sqlite db file: %s", err) 
        }
        file, err := os.Create(databasePath + "/grocery_store.db")
        if err != nil{
            return nil, fmt.Errorf("Error creating sqlite db file: %s", err) 
        }else{
            file.Close()
        }
        db, err := sql.Open("sqlite3", databasePath)
        if err != nil{
            fmt.Println("Error creating database:")
            fmt.Println(err)
            return nil, fmt.Errorf("There was an error creating the database: %s", err)
        }
        return db, nil
    }
}


func CreateDefaultTables(db *sql.DB, tableNames []string)(error){


    return nil
}
