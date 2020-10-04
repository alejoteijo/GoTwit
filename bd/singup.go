package bd

import (
	"context"
	"github.com/alejoteijo/GoTwit/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//AddUser insert user in database
func AddUser(user models.User) (string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("GoTwit")
	collection := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, error := collection.InsertOne(ctx, &user)
	if error != nil {
		return "", false, error
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil
}
