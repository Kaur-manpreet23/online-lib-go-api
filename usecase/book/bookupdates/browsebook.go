package bookupdates

import (
	"goborrow/domain" // importing structures
	"goborrow/repository/bookrepo"
	"goborrow/repository/logs" 
)

func BrowseBook(data domain.Data1)([]domain.Books, error) {
	
	search_val := data.Search_val
	search_category := data.Search_category
	search_data := data.Search_data
	if search_val == "View_all"{
		books,err := bookrepo.ViewAll()
		logs.InsertLog("View all books action invoked")
		return books,err
	}else{
		books,err := bookrepo.ViewBooks(search_category,search_data)
		logs.InsertLog("Viewing books as per"+search_category)
		return books,err
	}
}