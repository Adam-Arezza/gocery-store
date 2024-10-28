package database

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
    "embed"
    _ "github.com/mattn/go-sqlite3"
)


//go:embed migrations/create_tables.sql   
var migrations embed.FS

//go:embed migrations/default_data.sql
var defaultData embed.FS


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
        db, err := sql.Open("sqlite3", databasePath + "/grocery_store.db")

        if err != nil{
            return nil, fmt.Errorf("There was an error creating the database: %s", err)
        }

        err = createDefaultTables(db)
        if err != nil{
            return nil,err        
        }

        err = addDefaultData(db)
        if err != nil{
            return nil,err        
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

        db, err := sql.Open("sqlite3", databasePath + "/grocery_store.db")

        if err != nil{
            return nil, fmt.Errorf("There was an error creating the database: %s", err)
        }

        err = createDefaultTables(db)
        if err != nil{
            return nil,err        
        }

        err = addDefaultData(db)
        if err != nil{
            return nil,err        
        }

        return db, nil
    }
}


func createDefaultTables(db *sql.DB) error{
    migration, err := migrations.ReadFile("migrations/create_tables.sql")

    if err != nil{
        return fmt.Errorf("Error reading migration file: %s", err)
    }

    _, err = db.Exec(string(migration))

    if err != nil{
        return fmt.Errorf("Error creating tables: %s", err)
    }

    return nil
}


func checkDefaultData(tableName string, db *sql.DB) bool{
    var rowCount int
    rowQuery := fmt.Sprintf(`SELECT COUNT(*) from %s;`, tableName)
    rows, err := db.Query(rowQuery)
    if err != nil{
        fmt.Printf("error checking rows: %s", err)
    }
    for rows.Next(){
        if err:= rows.Scan(&rowCount); err != nil{
            fmt.Printf("error checking rows of %s: %s", tableName, err)
        }
    }

    if rowCount == 0{
        return false
    }
    return true
}


func addDefaultData(db *sql.DB) error{
    //check if the default data already exists
    if checkDefaultData("categories", db){
        return nil
    }

    dataMigration, err := defaultData.ReadFile("migrations/default_data.sql")

    if err != nil{
        return fmt.Errorf("Couldn't add default data: %s", err)
    }

    _, err = db.Exec(string(dataMigration))

    if err != nil{
        return fmt.Errorf("Couldn't add default data: %s", err)
    }

    return nil
}

