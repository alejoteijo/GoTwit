package bd

import (
	"github.com/alejoteijo/GoTwit/models"
	"golang.org/x/crypto/bcrypt"
)

func AttemptLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserExists(email)
	if found == false{
		return user, false
	}

	passwordDB := []byte(user.Password)
	passwordBytes := []byte(password)

	error := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if error != nil{
		return user, false
	}

	return user, true
}
