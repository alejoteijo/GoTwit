package middleware

import (
	"github.com/alejoteijo/GoTwit/bd"
	"net/http"
)

//CheckBD check BD connection
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if !bd.CheckConnection() {
			http.Error(response, "Database lost connection", 500)
			return
		}
		next.ServeHTTP(response, request)
	}
}
