package routers

import (
	"errors"
	"strings"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/alejoteijo/GoTwit/bd"
	"github.com/alejoteijo/GoTwit/models"
)

func processToken(token string) (*models.Claim, bool, string, error){
	privateKey := []byte(os.Getenv("TOKEN_KEY"))
}

