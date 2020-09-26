package routers

import (
	"encoding/json"
	"github.com/alejoteijo/GoTwit/bd"
	"github.com/alejoteijo/GoTwit/models"
	"net/http"
)

//SignUp user sign up
func SignUp(response http.ResponseWriter, request *http.Request) {
	var user models.User
	error := json.NewDecoder(request.Body).Decode(&user)
	if error != nil {
		http.Error(response, "Received data Error "+error.Error(), 400)
		return
	}

	//mail validation
	if len(user.Email) == 0 {
		http.Error(response, "Required Email", 400)
		return
	}

	//password validation
	if len(user.Password) < 6 {
		http.Error(response, "Password must be at least 6 characters", 400)
		return
	}

	_, found, _ := bd.CheckUserExists(user.Email)
	if found == true {
		http.Error(response, "User already exists in database", 400)
		return
	}

	_, status, error := bd.AddUser(user)
	if error != nil {
		http.Error(response, "Sign Up user error: "+error.Error(), 400)
		return
	}

	if status == false {
		http.Error(response, "Can't add user", 400)
		return
	}

	response.WriteHeader(http.StatusCreated)

}
