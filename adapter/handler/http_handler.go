package handler

import (
	"goborrow/domain"
	"goborrow/usecase/book/bookupdates"
	"goborrow/usecase/borrow/borrowupdates"
	"io"
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

func AddBook(w http.ResponseWriter, r *http.Request){
	body, err := ReturnBody(w,r)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
    }

	var book domain.Books
	err2 := json.Unmarshal([]byte(body), &book)
	if err2 != nil {
		fmt.Println(err2);
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	r_str, err1 := bookupdates.AddBook(book)
	if err1 != nil {
        http.Error(w, "Error adding book", http.StatusInternalServerError)
        return
    }
	responseData := map[string]string{"message": r_str}
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err3 := json.Marshal(responseData)
	if err3 != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}

func BrowseBook(w http.ResponseWriter, r *http.Request){
	body, err := ReturnBody(w,r)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
	}

	var data domain.Data1
	err2 := json.Unmarshal([]byte(body), &data)
	if err2 != nil {
		fmt.Println(err2);
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	var book []domain.Books;
	book,err1 := bookupdates.BrowseBook(data)
	if err1 != nil {
        http.Error(w, "Error browsing book", http.StatusInternalServerError)
        return
    }	
	j, err3 := json.Marshal(book)
	if err3 != nil {
        http.Error(w, "Failed to add marshal response", http.StatusInternalServerError)
        return
    }
	fmt.Println(string(j))
	w.Write(j)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	body, err := ReturnBody(w,r)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
	}
	var book domain.Books
	err2 := json.Unmarshal([]byte(body), &book)
	if err2 != nil {
		fmt.Println(err2);
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	res_str,err1 := bookupdates.DeleteBook(book)
	if err1 != nil {
        http.Error(w, "Error deleting book", http.StatusInternalServerError)
        return
    }
	responseData := map[string]string{"message": res_str}
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err3 := json.Marshal(responseData)
	if err3 != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}

func BorrowBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("Borrowing Books")
	body, err := ReturnBody(w,r)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		fmt.Printf("Failed to read request body\n")
        return
	}
	var data domain.Manage_borrows
	err2 := json.Unmarshal([]byte(body), &data)
	if err2 != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Println(err2)
		return
	}
	res_str,err1 := borrowupdates.BorrowBook(data)
	if err1 != nil {
		fmt.Printf("Error Borrowing Book")
        http.Error(w, "Error deleting book", http.StatusInternalServerError)
        return
    }	
	responseData := map[string]string{"message": res_str}
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err3 := json.Marshal(responseData)
	if err3 != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}


func ReturnBook(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Returning Books")
	body, err := ReturnBody(w,r)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
	}
	var data domain.Data2
	err2 := json.Unmarshal([]byte(body), &data)
	if err2 != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Println(err2)
		return
	}

	res_str,err1 := borrowupdates.ReturnBook(data)
	if err1 != nil {
        http.Error(w, "Error deleting book", http.StatusInternalServerError)
        return
    }	
	responseData := map[string]string{"message": res_str}
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err3 := json.Marshal(responseData)
	if err3 != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)	
}

func ReturnBody(w http.ResponseWriter, r *http.Request)([]byte, error){
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
	var body []byte
    for {
        buf := make([]byte, 1024)
        n, err := r.Body.Read(buf)
        if err != nil && err != io.EOF {
            http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Return Body error")
            return nil,err
        }
        if n == 0 {
            break
        }
        body = append(body, buf[:n]...)
    }
	return body,nil
}

func ValidateToken(token string)(bool,error){

	/*posturl := "http://localhost:3001/validate"
    
	resp, err := http.GET(posturl, "application/json"+token)
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return false, errors.New("token validation failed")
    }

    var isValid bool
    err = json.NewDecoder(resp.Body).Decode(&isValid)
    if err != nil {
        return false, err
    }

    return isValid, nil*/
	posturl := "http://localhost:3001/validate"

    req, err := http.NewRequest("GET", posturl, nil)
    if err != nil {
		fmt.Println(err)
        return false, err
    }

    req.Header.Set("Authorization", token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		fmt.Println(err)
        return false, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
        return false, errors.New(fmt.Sprintf("token validation failed with status: %d", resp.StatusCode))
    }

    var data domain.Data3
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
        return false, err
    }
	fmt.Println(data)
	fmt.Printf("user id while fetching is: %d and "+strconv.Itoa(data.UserId)+"+\n", data.UserId)
	os.Setenv("USERID", strconv.Itoa(data.UserId))

    return data.IsValid, nil
}
