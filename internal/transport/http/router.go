package http

import (
	"net/http"
	"taskgraph/internal/services"
	"taskgraph/internal/transport/http/handlers"
	"taskgraph/internal/transport/http/middleware"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(taskService services.TaskService) http.Handler {
	mux := http.NewServeMux()

	// Internal routes
	mux.HandleFunc("GET /health", handlers.HealthCheckHandler)

	// Swagger
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	// Task routes
	taskHandler := handlers.NewTaskHandler(taskService)
	mux.HandleFunc("POST /api/v1/tasks/create", taskHandler.Create)
	mux.HandleFunc("POST /api/v1/tasks/detail", taskHandler.Detail)

	// TODO: Create middleware chain function

	// Applying middlewares globally to the router
	handler := middleware.ResponseHeaderMiddleware(
		mux,
		"Content-Type",
		"application/json",
	)
	handler = middleware.LogMiddleware(handler)
	handler = middleware.ErrorHandleMiddleware(handler)

	return handler
}
