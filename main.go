package main

import (
	"context"
	"log"
	"net/http"

	"github.com/thalissonfelipe/ugly-api/database"
	"github.com/thalissonfelipe/ugly-api/routes"
)

func main() {
	client := database.InitializeMongoDB()
	defer client.Disconnect(context.Background())

	router := routes.NewRouter(client)

	log.Println("Server listening on port 5000!")
	log.Fatal(http.ListenAndServe(":5000", router))
}
