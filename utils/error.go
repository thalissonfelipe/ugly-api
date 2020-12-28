package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type httpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Path    string `json:"path"`
}

var httpStatusMessage = map[int]string{
	400: "Invalid JSON.",
	401: "Invalid Password.",
	404: "Resource not found.",
	409: "Resource already exists.",
	500: "Internal Error.",
}

var ErrAlreadyExists error = errors.New("mongo: resource already exists")
var ErrMissingRequiredFields error = errors.New("invalid format: missing required fields.")
var ErrInvalidUsernameLength error = errors.New("invalid format: username must be a minimum of 6 and a maximum of 20")
var ErrInvalidPasswordLength error = errors.New("invalid format: password must be a minimum of 8 and a maximum of 20")

// HandlerError - Global HandlerError function
func HandlerError(w http.ResponseWriter, r *http.Request, status int, err error) {
	if status == 0 {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			status = 401
			break
		case mongo.ErrNoDocuments:
			status = 404
			break
		case ErrAlreadyExists:
			status = 409
		default:
			status = 500
			break
		}
	}
	e := httpError{
		Message: httpStatusMessage[status],
		Status:  status,
		Error:   err.Error(),
		Path:    r.RequestURI,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(&e)
	CustomLogger.ErrorLogger.Println("-", err.Error())
}
