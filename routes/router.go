package routes

import (
	"github.com/gorilla/mux"
	"github.com/thalissonfelipe/ugly-api/controllers"
	"github.com/thalissonfelipe/ugly-api/middlewares"
	"github.com/thalissonfelipe/ugly-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewRouter is a custom function that creates a router with their specified routes and handlers
func NewRouter(client *mongo.Client) *mux.Router {
	mservice := services.MService{Client: client}
	mcontroller := controllers.MController{MService: &mservice}
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middlewares.LoggingMiddleware)

	api := router.PathPrefix("/api/v1/movies").Subrouter()
	api.HandleFunc("/", mcontroller.Index).Methods("GET")
	api.HandleFunc("/{name}", mcontroller.Show).Methods("GET")
	api.HandleFunc("/", mcontroller.Create).Methods("POST")
	api.HandleFunc("/{name}", mcontroller.Update).Methods("PUT")
	api.HandleFunc("/{name}", mcontroller.Delete).Methods("DELETE")

	return router
}
