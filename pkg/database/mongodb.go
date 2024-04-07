// database/mongodb.go

package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Client is the MongoDB client that will be initialized and used across the application
	Client *mongo.Client
	// DB is the database instance to be used across the application
	DB *mongo.Database
)

// Initialize sets up the MongoDB client and database instance
func Initialize() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Your MongoDB URI

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	DB = Client.Database("pokedex")
}
