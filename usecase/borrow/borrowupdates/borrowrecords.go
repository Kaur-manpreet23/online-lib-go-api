package borrowupdates

import (
	"fmt"
	"goborrow/domain" // importing structures
	"log"
	"goborrow/repository/logs" 
	"goborrow/repository/db" 
)

func BorrowBook(data domain.Manage_borrows)(string,error){
	
	bid := data.Bid
	uid := data.Uid
	
	//CHECKING AVAILABILITY OF BOOK
	selectquery:= "select status,quantity from books where id = ?;"
	rows, err3 := db.Db.Query(selectquery,bid)
	if err3 != nil {
		fmt.Printf("Query error")
		log.Fatal(err3)
		logs.InsertLog("Select Query error in BorrowBook action")
		return "",err3
	}	
	defer rows.Close()
	var status int
	var quantity int
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var b domain.Books
		if err := rows.Scan(&b.Status,&b.Quantity); err != nil {
			fmt.Printf("error")
			log.Fatal(err)
			logs.InsertLog("Scan error in Borrowbook action")
			return "",err 
		}
		status = b.Status
		quantity = b.Quantity
	}

	var resp_str string


	//UPDATING BORROW RECORDS
	if(status==0){
		resp_str = "Book Unavailable"
		} else {
		insertborrow := "insert into manage_borrows(bid, uid,issue_date,created_at,updated_at) values(?,?,(select curdate()),now(),now());"
		fmt.Printf(insertborrow)
		_, err := db.Db.Exec(insertborrow,bid,uid)
		if err != nil {
			log.Fatal(err)
			resp_str = "Manage Borrow Record Error"
		} else {

			var updatequery string
			if(quantity<=1){
				updatequery = "UPDATE books SET status=false,quantity=quantity-1 WHERE id = ?;"
			} else{
				updatequery = "UPDATE books SET quantity=quantity-1 WHERE id = ?;"
			}
			_, err := db.Db.Exec(updatequery, bid)
			if err != nil {
				fmt.Println(err);
				resp_str = "Status Update Error"
			} else {
				resp_str = "Borrow Success"
			}

		}
	}
	logs.InsertLog("Invoked BorrowBook action with status "+resp_str)
	return resp_str,nil
	
}