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
	mservice := services.MovieService{Client: client}
	uservice := services.UserService{Client: client}
	mhandler := handlers.MovieHandler{MovieService: &mservice}
	uhandler := handlers.UserHandler{UserService: &uservice}
	router := mux.NewRouter()

	router.Use(middlewares.LoggingMiddleware)

	jwtMiddleware := middlewares.JWTMiddleware(uservice.Client)

	// Movie routes
	router.HandleFunc("/api/v1/movies", jwtMiddleware(mhandler.ListMoviesHandler)).Methods("GET")
	router.HandleFunc("/api/v1/movies/{name}", jwtMiddleware(mhandler.GetMovieHandler)).Methods("GET")
	router.HandleFunc("/api/v1/movies", jwtMiddleware(mhandler.CreateMovieHandler)).Methods("POST")
	router.HandleFunc("/api/v1/movies/{name}", jwtMiddleware(mhandler.UpdateMovieHandler)).Methods("PUT")
	router.HandleFunc("/api/v1/movies/{name}", jwtMiddleware(mhandler.DeleteMovieHandler)).Methods("DELETE")

	// User routes
	router.HandleFunc("/api/v1/users", jwtMiddleware(uhandler.GetUsersHandler)).Methods("GET")
	router.HandleFunc("/api/v1/users", jwtMiddleware(uhandler.CreateUserHandler)).Methods("POST")
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
