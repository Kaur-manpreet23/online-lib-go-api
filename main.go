package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "goborrow/repository/db"
    "github.com/gorilla/handlers"
    "goborrow/adapter/authentication"
    "goborrow/adapter/handler"
)

func main() {
    db.Configure()
    
    r := mux.NewRouter()
    r.HandleFunc("/sbookadd", authentication.AuthenticateMiddleware(handler.AddBook)).Methods("POST")
    r.HandleFunc("/sbookbrowse", authentication.AuthenticateMiddleware(handler.BrowseBook)).Methods("POST")
    r.HandleFunc("/sbookdelete", authentication.AuthenticateMiddleware(handler.DeleteBook)).Methods("POST")
    r.HandleFunc("/mbookborrow", authentication.AuthenticateMiddleware(handler.BorrowBook)).Methods("POST")
    r.HandleFunc("/mbookreturn", authentication.AuthenticateMiddleware(handler.ReturnBook)).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000","http://localhost:3001"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization","user"}),
        handlers.AllowCredentials(),
    )(r)))
}
