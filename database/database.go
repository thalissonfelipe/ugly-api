package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitializeMongoDB configutarions
func InitializeMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	const URI string = "mongodb://localhost:27017/ugly_db"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected with mongodb!")

	return client
}
