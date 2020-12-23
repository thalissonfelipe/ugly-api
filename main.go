package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie Model
type Movie struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	IMDB     float32   `json:"imdb"`
	Category string    `json:"category"`
	Synosis  string    `json:"synopsis"`
	Director *Director `json:"director"`
}

// Director Model
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Sex       string `json:"sex"`
	Age       int16  `json:"age"`
}

// Mock movies
var movies []Movie

// Get all movies
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Get a movie by ID if it exists
func show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Movie not found."))
}

// Create a new movie
func create(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	movies = append(movies, movie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

// Update a movie by id if it exists
func update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			var movieUpdate Movie
			err := json.NewDecoder(r.Body).Decode(&movieUpdate)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			// Ugly method
			movies = append(movies[:index], movies[index+1:]...)
			movieUpdate.ID = movie.ID
			movies = append(movies, movieUpdate)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Movie not found."))
}

// Delete a movie by id if it exists
func delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Movie not found."))
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	// Mock data
	movies = append(movies, Movie{
		ID:       "1",
		Name:     "Harry Potter",
		IMDB:     7.5,
		Category: "Fantasy",
		Synosis:  "Lorem loren lorem lorem lorem...",
		Director: &Director{
			Firstname: "J",
			Lastname:  "K",
			Sex:       "Female",
			Age:       60,
		},
	})

	api := r.PathPrefix("/api/v1/movies").Subrouter()
	api.HandleFunc("/", index).Methods("GET")
	api.HandleFunc("/{id}", show).Methods("GET")
	api.HandleFunc("/", create).Methods("POST")
	api.HandleFunc("/{id}", update).Methods("PUT")
	api.HandleFunc("/{id}", delete).Methods("DELETE")

	fmt.Println("Server listening on port 5000!")

	log.Fatal(http.ListenAndServe(":5000", r))
}
