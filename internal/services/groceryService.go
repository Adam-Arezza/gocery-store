package services

import(
    "database/sql"
    "log"
    "fmt"
    "github.com/Adam-Arezza/gocery-store/internal/models"
)

func GetGroceryItems(db *sql.DB)([]models.GroceryItem, error){
    var groceries []models.GroceryItem
    groceryQuery := `SELECT * from grocery_items;`
    rows, err := db.Query(groceryQuery)

    if err != nil{
        log.Printf("Error getting grocery items: %s", err.Error())
        return nil, err
    }

    defer rows.Close()
    for rows.Next(){
        var groceryItem models.GroceryItem
        err := rows.Scan(&groceryItem.Id, &groceryItem.Name, &groceryItem.UnitPrice, &groceryItem.Stock, &groceryItem.CategoryId);
        if err != nil{
            log.Printf("Error reading database: %s", err.Error())
            return nil, err             }
        groceries = append(groceries, groceryItem)
    }
    return groceries, nil
}

func GetGroceryItemById(id int, db *sql.DB)(*models.GroceryItem, error){
    var groceryItem models.GroceryItem
    groceryQuery := `SELECT * from grocery_items WHERE id = ?;`
    row := db.QueryRow(groceryQuery, id).Scan(&groceryItem.Id,&groceryItem.Name,&groceryItem.UnitPrice,&groceryItem.Stock,&groceryItem.CategoryId)
    if row == sql.ErrNoRows{
        return nil, fmt.Errorf("No item found with id %d", id)
    }
    return &groceryItem, nil
}


func GetGroceryItemByName(name string, db *sql.DB)(*models.GroceryItem, error){
    var groceryItem models.GroceryItem
    groceryQuery := fmt.Sprintf(`SELECT * FROM grocery_items WHERE name='%s'`, name)
    row := db.QueryRow(groceryQuery, name).Scan(&groceryItem.Id, 
                                                &groceryItem.Name, 
                                                &groceryItem.UnitPrice, 
                                                &groceryItem.Stock, 
                                                &groceryItem.CategoryId)
    log.Println(groceryItem)
    if row != sql.ErrNoRows{
        log.Println(groceryQuery)
        return &groceryItem, nil 
    }else{
        return nil, fmt.Errorf("Error getting grocery item: %s", name)

    }
}

func UpdateGroceryItem(db *sql.DB, itemId int, newStock int)error{
    updateQuery := `UPDATE grocery_items SET stock = ? WHERE id = ?;`
    result, err := db.Exec(updateQuery, newStock, itemId)
    if err!= nil{
        return fmt.Errorf("Error updating stock items: %s", err.Error())
    }
    rowsAffected,_ := result.RowsAffected()
    if rowsAffected != 0 {
        fmt.Printf("updated %d rows\n", rowsAffected)
    }
    return nil
}
