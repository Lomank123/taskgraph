package domain

import (
	"time"

	"gorm.io/gorm"
)

// TaskDep represents the dependency between two tasks
type TaskDep struct {
	TaskID    string         `json:"task_id" gorm:"primaryKey"`
	DepID     string         `json:"dep_id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
