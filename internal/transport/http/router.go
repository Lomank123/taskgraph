package http

import (
	"net/http"
	"taskgraph/internal/services"
	"taskgraph/internal/transport/http/handlers"
	"taskgraph/internal/transport/http/middleware"
)

func NewRouter(taskService services.TaskService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HealthCheckHandler)

	taskHandler := handlers.NewTaskHandler(taskService)
	mux.HandleFunc("POST /task", taskHandler.Create)
	mux.HandleFunc("GET /task/{id}", taskHandler.GetByID)

	// TODO: Apply middleware chain function
	// Applying middlewares globally to the router
	return middleware.LogMiddleware(
		middleware.ResponseHeaderMiddleware(
			mux,
			"Content-Type",
			"application/json",
		),
	)
}
