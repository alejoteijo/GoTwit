package middleware

import (
	"github.com/alejoteijo/GoTwit/routers"
	"net/http"
)

func CheckJWT(next http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		_, _, _, error := routers.ProcessToken(r.Header.Get("Authorization"))
		if error != nil{
			http.Error(w, "Token Error: " + error.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w,r)
	}
}

