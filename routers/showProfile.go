package routers

import (
	"encoding/json"
	"net/http"
	"github.com/alejoteijo/GoTwit/bd"
)

func showProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "Should send ID parameter", http.StatusBadRequest)
		return
	}

	profile, error := bd.SearchProfile(ID)
	if error != nil{
		http.Error(w, "Error trying find profile"+ error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
