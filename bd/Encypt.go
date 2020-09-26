package bd

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error){
	difficulty := 6 //number of iterations to encrypt the password
	bytes, error := bcrypt.GenerateFromPassword([]byte(pass), difficulty)
	return string(bytes), error
}