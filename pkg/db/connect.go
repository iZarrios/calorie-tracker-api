package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var Client *mongo.Client = DBInstance()
var MONGO_URI string

func DBInstance() *mongo.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load .env file")
	}
	MONGO_URI = os.Getenv("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\nconnected to db\n\n")

	return client
}

func OpenCollection(client *mongo.Client, collectionNme string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("caloriesdb").Collection(collectionNme)

	return collection

}
