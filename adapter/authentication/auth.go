package authentication

import (
	"net/http"
	//"goborrow/domain"
	//"goborrow/repository/db" 
	"fmt"
	"goborrow/adapter/handler"
)

func AuthenticateMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        authenticated,err := checkAuthentication(r)
        
        if (!authenticated || err!=nil) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    }
}
//for authentication send api call to rails 
func checkAuthentication(r *http.Request) (bool,error) {
	token := r.Header.Get("Authorization")

	fmt.Println(token)
	return handler.ValidateToken(token)
}