package domain

import (
	"taskgraph/pkg/client/enum"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        string          `json:"id" gorm:"primarykey;default:gen_random_uuid()"`
	Status    enum.TaskStatus `json:"status"`
	Type      enum.TaskType   `json:"type"`
	Retries   int             `json:"retries"`
	Error     map[string]any  `json:"error" gorm:"serializer:json"`
	Result    map[string]any  `json:"result" gorm:"serializer:json"`
	Payload   map[string]any  `json:"payload" gorm:"serializer:json"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at"`
}
