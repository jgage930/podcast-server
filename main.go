package main

import (
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	db := Connect()
	SetupDb(db)

	router := feedRouter(&FeedHandler{db: db})

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	server.ListenAndServe()
}

func Connect() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open("data.db"),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Failed to establish a connection to the database.")
	}

	return db
}

func SetupDb(db *gorm.DB) {
	db.AutoMigrate(&Feed{})
}

type FeedHandler struct {
	db *gorm.DB
}

func feedRouter(h *FeedHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", h.CreateFeed)

	return mux
}

type Feed struct {
	gorm.Model

	Name string `json:"name"`
	Url  string `json:"url"`
}

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
