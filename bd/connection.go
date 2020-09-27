package bd

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoConnection is the received MongoDB client
var MongoConnection = ConnectBD()
var clientOptions = options.Client().ApplyURI(os.Getenv("MONGO_URL"))

//ConnectBD Try to connect to a MongoDB
func ConnectBD() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, error := mongo.Connect(ctx, clientOptions)
	checkError(error)
	error = client.Ping(ctx, nil)
	checkError(error)
	log.Println("Connection Succesful")
	return client
}

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

//CheckConnection check if connection is active
func CheckConnection() bool {
	error := MongoConnection.Ping(context.TODO(), nil)
	return error == nil
}
