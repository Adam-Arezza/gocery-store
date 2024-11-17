package stockmanager

import (
    "database/sql"
    "log"
    //"github.com/Adam-Arezza/gocery-store/internal/handlers"
)

type UpdateItem struct {
    grocery_item_id int
    stock int
}

var lowStockItems []UpdateItem

func CheckInventory(db *sql.DB){
    //check all the stock levels for all the grocery items
    //any stock level that is below a threshold should be marked for updating
    var item UpdateItem
    stockQuery := `SELECT id, stock FROM grocery_items WHERE stock < 5;`
    rows, err := db.Query(stockQuery)
    if err != nil{
        log.Printf("Could not get items to update stock")
        return
    }
    for rows.Next(){
        rows.Scan(&item.grocery_item_id, &item.stock)
        lowStockItems = append(lowStockItems, item)
    }
}

func UpdateStock(id,stock int, db *sql.DB){ 
    var rows int64

    for i,v := range lowStockItems{

    }
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


func ProcessOrders(db *sql.DB){
//get all the orders that are in the created status
//remove the quantity in the order from the stock
//change order status to completed
}
