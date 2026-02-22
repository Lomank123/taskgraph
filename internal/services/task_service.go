package services

import (
	"log"
	"taskgraph/internal/constants"
	"taskgraph/internal/domain"
	"taskgraph/pkg/client/enum"

	"gorm.io/gorm"
)

type TaskService interface {
	Create(taskType enum.TaskType, payload map[string]any, status *enum.TaskStatus, retries *int, result *map[string]any, errorString *map[string]any) (*domain.Task, error)
	Detail(id string) (*domain.Task, error)
}

type taskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) TaskService {
	return taskService{db: db}
}

func (s taskService) Create(taskType enum.TaskType, payload map[string]any, status *enum.TaskStatus, retries *int, result *map[string]any, errorString *map[string]any) (*domain.Task, error) {
	taskStatus := enum.TaskStatusPending
	if status != nil {
		taskStatus = *status
	}

	taskRetries := constants.DefaultTaskRetries
	if retries != nil {
		taskRetries = *retries
	}

	taskResult := map[string]any{}
	if result != nil {
		taskResult = *result
	}

	taskErrorString := map[string]any{}
	if errorString != nil {
		taskErrorString = *errorString
	}

	task := &domain.Task{
		Payload: payload,
		Status:  taskStatus,
		Type:    taskType,
		Retries: taskRetries,
		Result:  taskResult,
		Error:   taskErrorString,
	}

	taskCreateResult := s.db.Create(task)
	if taskCreateResult.Error != nil {
		return nil, taskCreateResult.Error
	}
	log.Printf("Task created with ID: %s", task.ID)

	return task, nil
}

func (s taskService) Detail(id string) (*domain.Task, error) {
	var task domain.Task
	getTaskResult := s.db.First(&task, "id = ?", id)
	if getTaskResult.Error != nil {
		return nil, getTaskResult.Error
	}
	log.Printf("Task found with ID: %s", task.ID)
	return &task, nil
}
