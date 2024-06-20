package bookrepo
import (
	"fmt"
	"log"
	"goborrow/repository/db"
	"goborrow/domain"
	//"strings"
)

func AddingBooks(title string,author string, publication_date string, genre string, status int, quantity int)(string){

_, err := db.Db.Exec("INSERT INTO books (title, author, publication_date, genre, status, quantity, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())", 
			title, author, publication_date, genre, status,quantity)
	var r_str string
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Add failure")
		r_str = "Add failure"
	} else {
		fmt.Printf("Add Success")
		r_str = "Add Success"
	}
	return r_str
}

func DeletingBooks(id int64)(string){

	deleteQuery := "DELETE FROM books WHERE id = ?;"

var err_code string
fmt.Printf(deleteQuery)
	_, err := db.Db.Exec(deleteQuery,id)
	if err != nil {
		log.Fatal(err)
		//templates.ExecuteTemplate(w,"book.html","Add Failure")
		err_code = "DELETE_FAILURE"
	} else {
		//	templates.ExecuteTemplate(w,"book.html","Add Success")
		err_code = "DELETE_SUCCESS"
	}
	return err_code
}

func ViewAll()([]domain.Books,error){

	var books []domain.Books
	selectQuery := "SELECT id,title,author,publication_date,genre,status,quantity FROM books;"
	rows, err := db.Db.Query(selectQuery)
	//rows, err := db.Db.Query("SELECT * FROM books WHERE ? = '?'",search_category,search_data)
	if err != nil {
		fmt.Printf("Query error")
		log.Fatal(err)
		return nil,err
	}	
	defer rows.Close()
	var date []uint8
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var b domain.Books
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &date, &b.Genre, &b.Status, &b.Quantity); err != nil {
			fmt.Printf("error")
			log.Fatal(err)
			return nil,err
		}
		b.Publication_date = string(date)
		fmt.Printf("Title is: %d\n",b.Status)
		books = append(books, b)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil,err
	}
	return books,nil
}

func ViewBooks(search_category string, search_data string)([]domain.Books,error){

	var books []domain.Books
	selectQuery := "SELECT id,title,author,publication_date,genre,status,quantity FROM books WHERE " + search_category + " = '" + search_data + "' ;" 
	rows, err := db.Db.Query(selectQuery)
	if err != nil {
		fmt.Printf("Query error")
		log.Fatal(err)
		return nil,err
	}	
	defer rows.Close()
	var date []uint8
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var b domain.Books
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &date, &b.Genre, &b.Status, &b.Quantity); err != nil {
			fmt.Printf("error")
			log.Fatal(err)
			return nil,err
		}
		b.Publication_date = string(date)
		fmt.Printf("Title is: %d\n",b.Status)
		books = append(books, b)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil,err
	}
	return books,nil
}