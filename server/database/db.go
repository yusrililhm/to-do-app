package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const Uri = "mongodb://localhost:27017"

func ConnectMongo() (*mongo.Database, error) {
	// setup configuration
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(Uri))
	if err != nil {
		log.Fatal(err.Error())
	}

	// PING the MongoDB
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err.Error())
	}

	return client.Database("todo"), nil
}