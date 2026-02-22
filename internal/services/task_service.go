package services

import (
	"log"
	"taskgraph/internal/domain"
	"taskgraph/pkg/client/enum"

	"gorm.io/gorm"
)

type TaskService interface {
	Create(payload string) (*domain.Task, error)
	GetByID(id string) (*domain.Task, error)
}

type taskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) TaskService {
	return taskService{db: db}
}

func (s taskService) Create(payload string) (*domain.Task, error) {
	// TODO: change input params & Validate "payload" for valid JSON
	task := &domain.Task{
		Payload: payload,
		Status:  enum.TaskStatusPending,
		Type:    enum.TaskTypeHttp,
	}

	result := s.db.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Task created with ID: %s", task.ID)

	return task, nil
}

func (s taskService) GetByID(id string) (*domain.Task, error) {
	var task domain.Task
	// TODO: Handle error
	s.db.First(&task, "id = ?", id)
	log.Printf("Task found with ID: %s", task.ID)
	return &task, nil
}
