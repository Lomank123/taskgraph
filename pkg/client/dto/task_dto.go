package dto

import (
	"taskgraph/pkg/client/enum"
	"time"
)

type TaskCreateRequestDTO struct {
	Payload string        `json:"payload"`
	Type    enum.TaskType `json:"type"`
}

type TaskCreateResponseDTO struct {
	ID     string          `json:"id"`
	Status enum.TaskStatus `json:"status"`
	Type   enum.TaskType   `json:"type"`
}

type TaskGetByIDRequestDTO struct {
	ID string `json:"id"`
}

type TaskGetByIDResponseDTO struct {
	ID        string          `json:"id"`
	Status    enum.TaskStatus `json:"status"`
	Type      enum.TaskType   `json:"type"`
	Retries   int             `json:"retries"`
	Error     string          `json:"error"`
	Result    string          `json:"result"`
	Payload   string          `json:"payload"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
