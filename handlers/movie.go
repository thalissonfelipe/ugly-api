package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thalissonfelipe/ugly-api/models"
	"github.com/thalissonfelipe/ugly-api/services"
	"github.com/thalissonfelipe/ugly-api/utils"
)

// MovieHandler struct
type MovieHandler struct {
	MService *services.MService
}

// ListMoviesHandler returns a list of users
func (m *MovieHandler) ListMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := m.MService.GetMovies()
	if err != nil {
		utils.HandlerError(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

// GetMovieHandler returns a movie object
func (m *MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := m.MService.GetMovie(params["name"])
	if err != nil {
		utils.HandlerError(w, r, 0, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

// CreateMovieHandler adds a new movie
func (m *MovieHandler) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err)
		return
	}

	err = m.MService.CreateMovie(&movie)
	if err != nil {
		utils.HandlerError(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// UpdateMovieHandler updates a movie object
func (m *MovieHandler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Refactor
	// params := mux.Vars(r)
	// var movie models.Movie
	// err := json.NewDecoder(r.Body).Decode(&movie)
	// if err != nil {
	// 	utils.HandlerError(w, r, http.StatusBadRequest, err)
	// 	return
	// }

	// err = m.MService.UpdateMovie(params["name"], &movie)
	// if err != nil {
	// 	utils.HandlerError(w, r, 0, err)
	// 	return
	// }

	w.WriteHeader(http.StatusNoContent)
}

// DeleteMovieHandler deletes a movie
func (m *MovieHandler) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := m.MService.DeleteMovie(params["name"])
	if err != nil {
		utils.HandlerError(w, r, 0, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
