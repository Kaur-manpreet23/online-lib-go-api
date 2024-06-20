package borrowupdates

import (
	"fmt"
	"goborrow/domain" // importing structures
	"log"
	"database/sql"
	"goborrow/repository/db"
	"goborrow/repository/logs" 
)

func ReturnBook(data domain.Data2)(string,error) {
	borrow_id := data.Borrow_id
	bid := data.Bid
	uid := data.Uid
	fmt.Printf("borrow_id : %d bid: %d uid: %d\n",borrow_id,bid,uid)
	var x int 
	var resp_str string
	validateQuery := "select 1 from manage_borrows where id = ? and uid = ? and bid = ? and return_date is NULL;"
	
	row := db.Db.QueryRow(validateQuery,borrow_id,uid,bid)

    	if err := row.Scan(&x); err != nil {
        if err == sql.ErrNoRows {
			fmt.Printf("No such books in your records")
            resp_str = "No such books in your records"
        }else{
			resp_str = "Error retriving data"
		}
    } else{
	fmt.Printf(resp_str)
    updatequery1 := "update books set quantity = quantity + 1, status = true where id = ?;"
    updatequery2 := "update manage_borrows set return_date = (select curdate()) where id = ? ;"
    _, err := db.Db.Exec(updatequery1,bid)
    if err != nil {
	log.Fatal(err)
	resp_str = "Error Updating book"
	fmt.Println(err)
    } else {
    _, err := db.Db.Exec(updatequery2,borrow_id)
    if err != nil {
	log.Fatal(err)
	resp_str = "Error Updating Manage Borrow"
    } else {
    	resp_str = "Return Success"
    }
    }
}
	logs.InsertLog("ReturnBook action invoked with status "+resp_str)
	return resp_str,nil
}