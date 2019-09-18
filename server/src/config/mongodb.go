package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

func init() {
	log.Println("Connecting to MongoDB...")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	MongoDB, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("ERROR: Could not make the connection to MongoDB!!", err)
		return
	}

	// Check the connection
	err = MongoDB.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("ERROR: Could not ping MongoDB", err)
	}

	log.Println("Connected to MongoDB!")
}
