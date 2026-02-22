package handlers

import (
	"encoding/json"
	"net/http"
	"taskgraph/internal/services"
	"taskgraph/pkg/client/dto"
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
	// Decode request body
	var payload dto.TaskCreateRequestDTO
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	// TODO: Add input validation step

	// Execute service method
	task, err := h.service.Create("dummy payload")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	// Build result payload and encode it into response
	data := dto.TaskCreateResponseDTO{
		ID:     task.ID,
		Status: task.Status,
		Type:   task.Type,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func (h taskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Get task id from request

	// TODO: Add input validation step

	task, err := h.service.GetByID("f67d8b0f-d13d-4b50-8e67-ee2ed37b3dd8")
	if err != nil {
		// TODO: ???
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := dto.TaskGetByIDResponseDTO{
		ID:        task.ID,
		Status:    task.Status,
		Type:      task.Type,
		Retries:   task.Retries,
		Error:     task.Error,
		Result:    task.Result,
		Payload:   task.Payload,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	json.NewEncoder(w).Encode(data)
}
