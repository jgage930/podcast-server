package filestore

import (
	"gorm.io/gorm"
)

type TaskStatus string

const (
	Queued  = "Queued"
	Running = "Running"
	Success = "Success"
	Failed  = "Failed"
)

type Task struct {
	gorm.Model

	ID     uint       `gorm:"primaryKey" json:"id"`
	Url    string     `json:"url"`
	Status TaskStatus `json:"status"`
}

type TaskHandler struct {
	db *gorm.DB
}
