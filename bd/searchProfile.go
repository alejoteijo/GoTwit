package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/alejoteijo/GoTwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database("GoTwit")
	collection := db.Collection("users")

	var profile models.User
	objectID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objectID}

	error := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password=""
	if error != nil{
		return profile, error
	}

	return profile, nil
}