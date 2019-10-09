package util

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(GetServerUrl())
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetServerUrl() string {
	url := os.Getenv("MONGODB_URI")
	if url != "" {
		return url
	}
	return "mongodb://localhost:27017" // Local default
}

func GetDBName() string {
	name := os.Getenv("DBNAME")
	if name != "" {
		return name
	}
	return "uberorders" // Default
}
