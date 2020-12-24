package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thalissonfelipe/ugly-api/models"
	"github.com/thalissonfelipe/ugly-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// MController struct
type MController struct {
	MService *services.MService
}

// Index is a handler function that returns all movies
func (c *MController) Index(w http.ResponseWriter, r *http.Request) {
	movies, err := c.MService.GetMovies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

// Show is a handler function that returns a movie by name if it exists
func (c *MController) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := c.MService.GetMovie(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error."))
		return
	}

	if movie == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Movie not found."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

// Create is a handler function that creates a new movie inside the database
func (c *MController) Create(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	// defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON."))
		return
	}

	err = c.MService.CreateMovie(&movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// Update is a handler function that updates a movie by name if it exists
func (c *MController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	// defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON."))
		return
	}

	movie.Name = params["name"]
	err = c.MService.UpdateMovie(&movie)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Movie not found."))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error."))
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete is a handler function that deletes a movie by name if it exists
func (c *MController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := c.MService.DeleteMovie(params["name"])
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Movie not found."))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error."))
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
