package bookupdates

import (
	"goborrow/domain" // importing structures
	"goborrow/repository/bookrepo"
	"goborrow/repository/logs" 
)

func DeleteBook(book domain.Books)(string,error) {

	id := book.ID
	
	err_code := bookrepo.DeletingBooks(id)
	logs.InsertLog("Delete book action invoked with status "+err_code)
	return err_code,nil
}