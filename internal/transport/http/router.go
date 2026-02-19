package http

import (
	"net/http"
	"taskgraph/internal/services"
	"taskgraph/internal/transport/http/handlers"
)

func SetupRouter(taskService services.TaskService) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheckHandler)

	taskHandler := handlers.NewTaskHandler(taskService)
	mux.HandleFunc("/task", taskHandler.Create)
	mux.HandleFunc("/task/{id}", taskHandler.GetByID)
	return mux
}
