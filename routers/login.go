package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alejoteijo/GoTwit/bd"
	"github.com/alejoteijo/GoTwit/jwt"
	"github.com/alejoteijo/GoTwit/models"
)

func Login(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")

	var user models.User

	error := json.NewDecoder(request.Body).Decode(&user)

	if error != nil{
		http.Error(response, "Invalid user or password " + error.Error(),http.StatusBadRequest)
		return
	}

	if len(user.Email)==0{
		http.Error(response, "User email required ", http.StatusBadRequest)
		return
	}

	document, exists := bd.AttemptLogin(user.Email, user.Password)
	if !exists {
		http.Error(response, "Invalid user or password.", http.StatusBadRequest)
		return
	}

	jwtKey, error := jwt.GenerateJWT(document)
	if error != nil{
		http.Error(response, "Error generating Token "+error.Error(), http.StatusBadRequest)
		return
	}

	resp := models.LoginResponse{Token: jwtKey}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(response, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}

