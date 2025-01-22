package filestore

import (
	"net/http"

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
	db    *gorm.DB
	queue Queue
}

func TaskRouter(h *TaskHandler, mux *http.ServeMux) {
	mux.HandleFunc("GET /task/", nil)
	mux.HandleFunc("POST /task/push", nil)
	mux.HandleFunc("GET /task/status/{}", nil)
}

type TaskCreate struct {
	Url    string     `json:"url"`
	Status TaskStatus `json:"status"`
}

func (h *TaskHandler) pushTask(w http.ResponseWriter, r *http.Request) {
}
