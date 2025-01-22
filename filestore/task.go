package filestore

import (
	"net/http"
	"podcast-server/common"

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

	Url    string     `json:"url"`
	Status TaskStatus `json:"status"`
}

func NewTask(url string) Task {
	return Task{
		Url:    url,
		Status: "Queued",
	}
}

type TaskHandler struct {
	db    *gorm.DB
	queue Queue
}

func TaskRouter(h *TaskHandler, mux *http.ServeMux) {
	//mux.HandleFunc("GET /task/", nil)
	mux.HandleFunc("POST /task/push", h.pushTask)
	//mux.HandleFunc("GET /task/status/{}", nil)
}

type TaskCreate struct {
	Url string `json:"url"`
}

func (h *TaskHandler) pushTask(w http.ResponseWriter, r *http.Request) {
	var payload TaskCreate
	common.ReadBody(&payload, w, r)

	task := NewTask(payload.Url)
	h.db.Create(&task)

	h.queue.Push(task)

	common.Respond(task, w)
}
