package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"taskgraph/utils"
)

// ErrorHandleMiddleware is a middleware that recovers from panics and returns a 500 error.
func ErrorHandleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC RECOVERED: %v\n%s", err, debug.Stack())	
				utils.HandleAPIError(w, http.StatusInternalServerError, "internal server error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
