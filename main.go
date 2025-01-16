package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	log.Printf("Starting Application.")

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Podcast Server v0.0.1"))
	})

	router := feedRouter(&FeedHandler{})

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	server.ListenAndServe()
}

func feedRouter(h *FeedHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", h.CreateFeed)

	return mux
}

type Feed struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type FeedHandler struct{}

func (h *FeedHandler) CreateFeed(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
	log.Printf("Recieved feed %s", feed.Name)
}
