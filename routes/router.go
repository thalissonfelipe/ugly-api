package routes

import (
	"github.com/gorilla/mux"
	"github.com/thalissonfelipe/ugly-api/handlers"
	"github.com/thalissonfelipe/ugly-api/middlewares"
	"github.com/thalissonfelipe/ugly-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewRouter is a custom function that creates a router with their specified routes and handlers
func NewRouter(client *mongo.Client) *mux.Router {
	mservice := services.MService{Client: client}
	mhandler := handlers.MovieHandler{MService: &mservice}
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middlewares.LoggingMiddleware)

	api := router.PathPrefix("/api/v1/movies").Subrouter()
	api.HandleFunc("/", mhandler.ListMoviesHandler).Methods("GET")
	api.HandleFunc("/{name}", mhandler.GetMovieHandler).Methods("GET")
	api.HandleFunc("/", mhandler.CreateMovieHandler).Methods("POST")
	api.HandleFunc("/{name}", mhandler.UpdateMovieHandler).Methods("PUT")
	api.HandleFunc("/{name}", mhandler.DeleteMovieHandler).Methods("DELETE")

	return router
}

// type HandlerFunc func(http.ResponseWriter, *http.Request)

// type Route struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	HandlerFunc HandlerFunc
// }

// var Routes []Route{
// 	Route{"ListMovies", "GET", "/", }
// }
