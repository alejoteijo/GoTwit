package bd

import (
	"context"
	"github.com/alejoteijo/GoTwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func CheckUserExists(email string) (models.User, bool, string){
	ctx , cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConnection.Database("goTwit")
	col := db.Collection("users")

	condition := bson.M{"email":email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex() //transform result's ID in Hexadecimal string

	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
