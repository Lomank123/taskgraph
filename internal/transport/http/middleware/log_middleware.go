package middleware

import (
	"log"
	"net/http"
)

// LogMiddleware logs the request method and URL path, and the response headers.
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Print full info about request including body, headers, etc.
		log.Println("Request: ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		// TODO: Print full info about response including body, status code, etc.
		log.Println("Response: ", w.Header())
	})
}
