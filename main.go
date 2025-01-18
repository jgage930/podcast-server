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
	log.Printf("Starting Application.")

	testMux := http.NewServeMux()
	testMux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Called GET /test"))
	})
	testMux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Called POST /test"))
	})

	userMux := http.NewServeMux()
	userMux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Called GET /user"))
	})
	userMux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Called POST /user"))
	})

	topMux := http.NewServeMux()
	topMux.Handle("/test", testMux)
	topMux.Handle("/user", userMux)

	port := ":8080"
	http.ListenAndServe(port, topMux)
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
