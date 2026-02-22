package handlers

import (
	"encoding/json"
	"net/http"
	"slices"
	"taskgraph/internal/services"
	"taskgraph/pkg/client/dto"
	"taskgraph/pkg/client/enum"
	"taskgraph/utils"

	"github.com/google/uuid"
)

type TaskHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	service services.TaskService
}

func NewTaskHandler(service services.TaskService) TaskHandler {
	return taskHandler{service: service}
}

// Create godoc
// @Summary Create a task with the given payload and type
// @Param task body dto.TaskCreateRequestDTO true "Task Create Request"
// @Success 201 {object} dto.TaskCreateResponseDTO
// @Failure 400 {object} map[string]string
// @Router /v1/tasks/create [post]
func (h taskHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var reqPayload dto.TaskCreateRequestDTO
	err := json.NewDecoder(r.Body).Decode(&reqPayload)
	if err != nil {
		utils.HandleAPIError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate input
	if reqPayload.Retries != nil && *reqPayload.Retries < 0 {
		utils.HandleAPIError(w, http.StatusBadRequest, "invalid retries")
		return
	}
	if reqPayload.Status != nil && !slices.Contains(enum.AllTaskStatuses, *reqPayload.Status) {
		utils.HandleAPIError(w, http.StatusBadRequest, "invalid task status")
		return
	}
	if !slices.Contains(enum.AllTaskTypes, reqPayload.Type) {
		utils.HandleAPIError(w, http.StatusBadRequest, "invalid task type")
		return
	}

	// Execute service method
	task, err := h.service.Create(reqPayload.Type, reqPayload.Payload, reqPayload.Status, reqPayload.Retries, reqPayload.Result, reqPayload.Error)
	if err != nil {
		utils.HandleAPIError(w, http.StatusInternalServerError, "error occurred while creating task")
		return
	}

	// Build result payload and encode it into response
	data := dto.TaskCreateResponseDTO{
		Task: dto.TaskDTO{
			ID:        task.ID,
			Status:    task.Status,
			Type:      task.Type,
			Retries:   task.Retries,
			Error:     task.Error,
			Result:    task.Result,
			Payload:   task.Payload,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		},
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

// Detail godoc
// @Summary Get details of a task by ID
// @Param task body dto.TaskGetByIDRequestDTO true "Task Detail Request"
// @Success 200 {object} dto.TaskGetByIDResponseDTO
// @Router /v1/tasks/detail [post]
func (h taskHandler) Detail(w http.ResponseWriter, r *http.Request) {
	var reqPayload dto.TaskGetByIDRequestDTO
	err := json.NewDecoder(r.Body).Decode(&reqPayload)
	if err != nil {
		utils.HandleAPIError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if reqPayload.ID == "" {
		utils.HandleAPIError(w, http.StatusBadRequest, "task ID is required")
		return
	}
	if err := uuid.Validate(reqPayload.ID); err != nil {
		utils.HandleAPIError(w, http.StatusBadRequest, "invalid task ID format")
		return
	}

	task, err := h.service.Detail(reqPayload.ID)
	if err != nil {
		utils.HandleAPIError(w, http.StatusInternalServerError, "error occurred while fetching task")
		return
	}

	data := dto.TaskGetByIDResponseDTO{
		Task: dto.TaskDTO{
			ID:        task.ID,
			Status:    task.Status,
			Type:      task.Type,
			Retries:   task.Retries,
			Error:     task.Error,
			Result:    task.Result,
			Payload:   task.Payload,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		},
	}

	json.NewEncoder(w).Encode(data)
}
