package handlers

import (
	"fmt"
	"net/http"
	"taskgraph/internal/services"
)

type TaskHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	service services.TaskService
}

func NewTaskHandler(service services.TaskService) TaskHandler {
	return taskHandler{service: service}
}

func (h taskHandler) Create(w http.ResponseWriter, r *http.Request) {
	task, err := h.service.Create("dummy payload")
	if err != nil {
		// TOOD: ???
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO: Return JSON response

	fmt.Fprintf(w, "Task created with ID: %s", task.ID)
}

func (h taskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	task, err := h.service.GetByID("f67d8b0f-d13d-4b50-8e67-ee2ed37b3dd8")
	if err != nil {
		// TODO: ???
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO: Return JSON response

	fmt.Fprintf(w, "Task found with ID: %s", task.ID)
}
