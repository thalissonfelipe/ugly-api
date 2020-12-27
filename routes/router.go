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
	uservice := services.UService{Client: client}
	mhandler := handlers.MovieHandler{MService: &mservice}
	uhandler := handlers.UserHandler{UService: &uservice}
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middlewares.LoggingMiddleware)

	// Movie routes
	router.HandleFunc("/api/v1/movies", middlewares.JWTMiddleware(mhandler.ListMoviesHandler)).Methods("GET")
	router.HandleFunc("/api/v1/movies/{name}", middlewares.JWTMiddleware(mhandler.GetMovieHandler)).Methods("GET")
	router.HandleFunc("/api/v1/movies", middlewares.JWTMiddleware(mhandler.CreateMovieHandler)).Methods("POST")
	router.HandleFunc("/api/v1/movies/{name}", middlewares.JWTMiddleware(mhandler.UpdateMovieHandler)).Methods("PUT")
	router.HandleFunc("/api/v1/movies/{name}", middlewares.JWTMiddleware(mhandler.DeleteMovieHandler)).Methods("DELETE")

	// User routes
	router.HandleFunc("/api/v1/users", middlewares.JWTMiddleware(uhandler.GetUsersHandler)).Methods("GET")
	router.HandleFunc("/api/v1/users", middlewares.JWTMiddleware(uhandler.CreateUserHandler)).Methods("POST")
	router.HandleFunc("/api/v1/users/login", uhandler.LoginHandler).Methods("POST")

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
