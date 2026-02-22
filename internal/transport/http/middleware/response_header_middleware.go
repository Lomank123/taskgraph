package middleware

import "net/http"

// ResponseHeaderMiddleware sets the provided header to all responses.
func ResponseHeaderMiddleware(next http.Handler, headerName, headerValue string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerName, headerValue)
		next.ServeHTTP(w, r)
	})
}
