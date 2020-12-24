package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	c "github.com/thalissonfelipe/ugly-api/config"
	"github.com/thalissonfelipe/ugly-api/database"
	"github.com/thalissonfelipe/ugly-api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	c.MyConfig = c.GetConfig()

	client := database.InitializeMongoDB()
	defer client.Disconnect(context.Background())

	router := routes.NewRouter(client)

	port := c.MyConfig.API.Port
	if port == "" {
		port = "5000"
	}
	log.Printf("Server listening on port %s!", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
