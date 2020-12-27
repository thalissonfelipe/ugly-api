package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thalissonfelipe/ugly-api/models"
	"github.com/thalissonfelipe/ugly-api/services"
	"github.com/thalissonfelipe/ugly-api/utils"
)

// UserHandler struct
type UserHandler struct {
	UService *services.UService
}

// LoginHandler ...
func (u *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err)
		return
	}

	err = u.UService.Authenticate(&login)
	if err != nil {
		utils.HandlerError(w, r, 0, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := u.UService.GetUsers()
	if err != nil {
		utils.HandlerError(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (u *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err)
		return
	}

	err = u.UService.CreateUser(&user)
	if err != nil {
		utils.HandlerError(w, r, 0, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
