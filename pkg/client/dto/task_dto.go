package dto

import (
	"taskgraph/pkg/client/enum"
	"time"
)

type TaskDTO struct {
	ID        string          `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt time.Time       `json:"deleted_at"`
	Status    enum.TaskStatus `json:"status"`
	Type      enum.TaskType   `json:"type"`
	Retries   int             `json:"retries"`
	Result    map[string]any  `json:"result"`
	Payload   map[string]any  `json:"payload"`
	Error     map[string]any  `json:"error"`
}

type TaskCreateRequestDTO struct {
	Payload map[string]any   `json:"payload"`
	Type    enum.TaskType    `json:"type"`
	Status  *enum.TaskStatus `json:"status"`
	Retries *int             `json:"retries"`
	Result  *map[string]any  `json:"result"`
	Error   *map[string]any  `json:"error"`
}

type TaskCreateResponseDTO struct {
	Task TaskDTO `json:"task"`
}

type TaskGetByIDRequestDTO struct {
	ID string `json:"id"`
}

type TaskGetByIDResponseDTO struct {
	Task TaskDTO `json:"task"`
}
