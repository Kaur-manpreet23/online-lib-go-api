package bookupdates

import (
	"fmt"
	"goborrow/domain" // importing structures
	"goborrow/repository/db"
	"goborrow/repository/bookrepo"
	"log" 
	"goborrow/repository/logs" 
)

func AddBook(book domain.Books)(string, error){
	title := book.Title
	author := book.Author
	publication_date := book.Publication_date
	genre := book.Genre
	status := book.Status
	quantity := book.Quantity
	
	fmt.Printf("No value "+ title + author + publication_date + genre )

	pingErr := db.Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		logs.InsertLog("Db connection error")
		return "",pingErr
	}
	r_str := bookrepo.AddingBooks(title,author,publication_date,genre,status,quantity)
	logs.InsertLog("Addbooks action" + r_str)
	return r_str,nil
}