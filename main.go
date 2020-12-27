package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/thalissonfelipe/ugly-api/config"
	"github.com/thalissonfelipe/ugly-api/database"
	"github.com/thalissonfelipe/ugly-api/routes"
	"github.com/thalissonfelipe/ugly-api/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	config.MyConfig = config.GetConfig()

	client := database.InitializeMongoDB()
	defer client.Disconnect(context.Background())

	router := routes.NewRouter(client)

	port := config.MyConfig.API.Port
	if port == "" {
		port = "5000"
	}
	utils.CustomLogger.InfoLogger.Printf("- Server listening on port %s!", port)
	utils.CustomLogger.ErrorLogger.Fatal(http.ListenAndServe(":"+port, router))
}
