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
		http.Error(response, "Invalid user or password " + error.Error(),400)
		return
	}

	if len(user.Email)==0{
		http.Error(response, "User email required ", 400)
		return
	}

	document, exists := bd.loginAttempt(user.Email, user.Password)
	if exists == false {
		http.Error(response, "Invalid user or password", 400)
		return
	}

	jwtKey, error := jwt.generateJWT(document)
	if error != nil{
		http.Error(response, "Error generating Token "+error.Error(),400)
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

