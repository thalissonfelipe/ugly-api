package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type httpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Path    string `json:"path"`
}

var httpStatusMessage = map[int]string{
	400: "Invalid JSON.",
	404: "Resource not found.",
	500: "Internal Error.",
}

// HandlerError - Global HandlerError function
func HandlerError(w http.ResponseWriter, r *http.Request, status int, err string) {
	e := httpError{
		Message: httpStatusMessage[status],
		Status:  status,
		Error:   err,
		Path:    r.RequestURI,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(&e)
	errorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	errorLogger.Println(err)
}
