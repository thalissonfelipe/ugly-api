package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/thalissonfelipe/ugly-api/config"
	"github.com/thalissonfelipe/ugly-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitializeMongoDB configutarions
func InitializeMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db := config.MyConfig.DB
	var URI string

	if db.Username == "" || db.Password == "" {
		URI = fmt.Sprintf("mongodb://%s:%s/%s", db.Host, db.Port, db.DatabaseName)
	} else {
		URI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", db.Username, db.Password, db.Host, db.Port, db.DatabaseName)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	utils.CustomLogger.InfoLogger.Println("Connected with mongodb!")

	return client
}
