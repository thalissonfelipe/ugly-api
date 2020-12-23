package middlewares

import (
	"log"
	"net/http"
	"os"
)

// LoggingMiddleware is a function to log income requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
			infoLogger.Printf("- %s %s", r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
		})
}
