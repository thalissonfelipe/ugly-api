package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// ContextKey ...
type ContextKey string

// ContextKeyValue is a key to store each request ID
var ContextKeyValue ContextKey = "requestID"

// LogRecord is a struct to store the response http status
type LogRecord struct {
	http.ResponseWriter
	status int
}

func (r *LogRecord) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

// WriteHeader ...
func (r *LogRecord) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// LoggingMiddleware is a function to log income requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			record := &LogRecord{
				ResponseWriter: w,
			}

			ctx := r.Context()
			id := uuid.New().String()
			ctx = context.WithValue(ctx, ContextKeyValue, id)
			r = r.WithContext(ctx)

			infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
			infoLogger.Printf("- %s - %s %s", id, r.Method, r.RequestURI)

			next.ServeHTTP(record, r)

			infoLogger.Printf("- %s - Status Code %d", id, record.status)
		})
}

// GetRequestID is a function to get the request id
func GetRequestID(ctx context.Context) string {
	reqID := ctx.Value(ContextKeyValue)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}
