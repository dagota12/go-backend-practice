package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_URI string
var DB_NAME string

// NewMongoClient returns a new MongoDB client.
// It connects to the MongoDB instance specified by the DB_URI environment variable.
// If the connection fails, the function terminates the program.
func NewMongoClient() *mongo.Client {
	godotenv.Load()
	log.Println("Connecting to db...")
	DB_URI = os.Getenv("DB_URI")
	DB_NAME = os.Getenv("DB_NAME")

	clientOptions := options.Client().ApplyURI(DB_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected To MongoDB")

	return client
}
