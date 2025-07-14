package stockmanager

import (
	"database/sql"
	"log"
    "fmt"
	"github.com/Adam-Arezza/gocery-store/internal/handlers"
   // "github.com/Adam-Arezza/gocery-store/internal/models"   
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
    for{
        rows, err := db.Query(stockQuery)
        if err != nil{
            log.Printf("Could not get items to update stock")
            return
        }
        for rows.Next(){
            rows.Scan(&item.grocery_item_id, &item.stock)
            lowStockItems = append(lowStockItems, item)
            fmt.Printf("Item %d needs restocking\n", item.grocery_item_id)
        }
    }
}

func UpdateStock(db *sql.DB, newStock int){ 
    for i := range lowStockItems{
        handlers.UpdateGroceryItemHandler(db, lowStockItems[i].grocery_item_id, newStock) 
       // updateQuery := `UPDATE grocery_items SET stock = ? WHERE id = ?;`
       // result, err := db.Exec(updateQuery, 20, lowStockItems[i].grocery_item_id)
       // if err!= nil{
       //     fmt.Printf("Error updating stock items: %s", err.Error())
       // }
       // rowsAffected,_ := result.RowsAffected()
       // if rowsAffected != 0 {
       //     fmt.Printf("updated %d rows\n", rowsAffected)
       // }
    }
}

func ProcessOrders(db *sql.DB){
//get all the orders that are in the created status
//remove the quantity in the order from the stock
//change order status to completed
    orderQuery := `SELECT * from orders WHERE status_id = 1`
    currentOrders, err := db.Query(orderQuery)  
    if err != nil{
        fmt.Printf("Error in processing orders: %s", err.Error())
    }
    for currentOrders.Next(){
        //scan row to variable
    }
}

func ManageStock(db *sql.DB){
    CheckInventory(db)
    //UpdateStock()
}
