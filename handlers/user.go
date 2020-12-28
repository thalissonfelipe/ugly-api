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
	UserService *services.UserService
}

// LoginHandler ...
func (u *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err)
		return
	}

	err = u.UserService.Authenticate(&login)
	if err != nil {
		utils.HandlerError(w, r, 0, err)
		return
	}

	token, err := utils.CreateToken(login.Username, w)
	if err != nil {
		utils.HandlerError(w, r, 500, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Token{Token: token})
}

func (u *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserService.GetUsers()
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

	err = utils.ValidateCreateUserBody(user)
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err)
		return
	}

	err = u.UserService.CreateUser(&user)
	if err != nil {
		utils.HandlerError(w, r, 0, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
