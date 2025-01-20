package main

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type FeedHandler struct {
	db *gorm.DB
}

type Feed struct {
	gorm.Model

	Name string `json:"name"`
	Url  string `json:"url"`
}

func feedRouter(h *FeedHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.listFeeds)
	mux.HandleFunc("GET /{id}", h.getFeedById)
	mux.HandleFunc("POST /", h.createFeed)
	mux.HandleFunc("DELETE /{id}", h.deleteFeedById)

	return mux
}

func (h *FeedHandler) createFeed(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var feed Feed
	if err := json.Unmarshal(body, &feed); err != nil {
		http.Error(w, "Failed to parse json Body", http.StatusBadRequest)
		return
	}

	h.db.Create(&feed)

	response := Response{
		Message: "Successfully Created Feed",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *FeedHandler) getFeedById(w http.ResponseWriter, r *http.Request) {
	var feed Feed
	id := r.PathValue("id")
	GetById(&feed, id, h.db, w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feed)
}

func (h *FeedHandler) listFeeds(w http.ResponseWriter, r *http.Request) {
	var feeds []Feed
	h.db.Find(&feeds)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)
}

func (h *FeedHandler) deleteFeedById(w http.ResponseWriter, r *http.Request) {
	var feed Feed
	id := r.PathValue("id")
	err := h.db.First(&feed, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "Record Not Found", http.StatusNotFound)
		return
	}

	h.db.Delete(&feed)

	response := Response{Message: "Deleted Record"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
