package stockmanager

import (
    "database/sql"
    "log"
    //"github.com/Adam-Arezza/gocery-store/internal/handlers"
)

func CheckInventory(db *sql.DB){

}

func UpdateStock(id,stock int, db *sql.DB){ 
    var rows int64
    updateQuery := `UPDATE grocery_items SET stock = ? WHERE id = ?;`

    result, err := db.Exec(updateQuery, stock, id) 

    if err != nil{
        log.Printf("Error updating grocery item with id: %v", id)
        return
    }

    rows, err = result.RowsAffected()

    if err != nil{
        log.Printf("Error updating grocery item with id: %v \n%s", id, err.Error())
        return
    }

    if rows == 0{
        log.Printf("No item found")
    }
}
