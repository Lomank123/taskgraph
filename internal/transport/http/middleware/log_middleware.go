package middleware

import (
	"log"
	"net/http"
)

// LogMiddleware logs the request method and URL path, and the response headers.
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("Response: ", w.Header())
	})
}
