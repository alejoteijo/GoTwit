package middleware

import (
	"net/http"
	"github.com/alejoteijo/GoTwit/routers"
)

func checkJWT(next http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		_, _, _, error := routers.ProcessToken(r.Header.Get("Authorization"))
		if error != nil{
			http.Error(w, "Token Error: " + error.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w,r)
	}
}

