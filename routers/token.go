package routers

import (
	"errors"
	"os"
	"strings"

	"github.com/alejoteijo/GoTwit/bd"
	"github.com/alejoteijo/GoTwit/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var UserID string

func processToken(token string) (*models.Claim, bool, string, error){
	privateKey := []byte(os.Getenv("TOKEN_KEY"))
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) !=2 {
		return claims, false, "", errors.New("Invalid Token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, error := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token)(interface{},error){
		return privateKey, nil
	})
	if error==nil{
		_, found, _ := bd.CheckUserExists(claims.Email)
		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}
	if !tkn.Valid{
		return claims, false, "", errors.New("Invalid Token")
	}

	return claims, false, "", error
}

