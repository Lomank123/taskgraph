package http

import (
	"net/http"
	"taskgraph/internal/handlers"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheckHandler)
	return mux
}
