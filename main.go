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

	feedMux := feedRouter(&FeedHandler{db: db})

	otherMux := http.NewServeMux()
	otherMux.HandleFunc("GET /other", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Called other endpoint"))
	})

	defaultMux := http.NewServeMux()
	defaultMux.Handle("/feed", http.StripPrefix("/feed", feedMux))
	defaultMux.Handle("/other", otherMux)

	server := &http.Server{
		Addr:    port,
		Handler: defaultMux,
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

type Response struct {
	Message string `json:"Message"`
}

type FeedHandler struct {
	db *gorm.DB
}

func feedRouter(h *FeedHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.CreateFeed)
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

	h.db.Create(&feed)

	response := Response{
		Message: "Successfully Created Feed",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *FeedHandler) ListFeeds(w http.ResponseWriter, r *http.Request) {
	log.Println("got endpoint")
}
