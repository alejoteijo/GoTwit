package routers

import (
	"../bd"
	"../models"
	"encoding/json"
	"net/http"
)

//SignUp user sign up
func SignUp(response http.ResponseWriter, request *http.Request) {
	var t models.User
	error := json.NewDecoder(request.Body).Decode(&t)
	if error != nil {
		http.Error(response, "Received data Error "+error.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(response, "Required Email", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(response, "Password must be at least 6 characters", 400)
		return
	}

	_, finded, _ := bd.CheckUserExists(t.Email)
	if finded == true {
		http.Error(response, "User already exists in database", 400)
		return
	}

	_, status, error := bd.AddRegister(t)
	if error != nil {
		http.Error(response, "Sign Up user error "+error.Error(), 400)
		return
	}

	if status == false {
		http.Error(response, "Can't add user", 400)
		return
	}

	response.WriteHeader(http.StatusCreated)

}
