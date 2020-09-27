package jwt

import (
	"os"
	"time"

	"github.com/alejoteijo/GoTwit/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func generateJWT(user models.User) (string, error) {

	privateKey := []byte(os.Getenv("TOKEN_KEY"))

	payload := jwt.MapClaims{
		"email":      user.Email,
		"name":       user.Name,
		"lastName":   user.LastName,
		"birthDate":  user.BirthDate,
		"biography":  user.Biography,
		"Location":   user.Location,
		"website":    user.Website,
		"_id":        user.ID.Hex(),
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, error := token.SignedString(privateKey)
	if error != nil {
		return tokenString, error
	}

	return tokenString, nil
}
