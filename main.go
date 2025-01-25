package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"podcast-server/filestore"
	"strconv"
)

func NewServer() *http.ServeMux {
	log.Println("Creating Server Mux")
	mux := http.NewServeMux()
	return mux
}

func main() {
	log.Printf("Starting Application.")

	db := Connect()
	SetupDb(db)

	// Create New Server
	router := http.NewServeMux()

	// Add routes
	feedHandler := FeedHandler{db: db}
	FeedRouter(&feedHandler, router)

	podcastHandler := PodcastHandler{db: db}
	PodcastRouter(&podcastHandler, router)

	taskHandler := filestore.NewTaskHandler(db)
	filestore.TaskRouter(&taskHandler, router)

	// Set up top level middleware
	configured := LoggingMiddleware(router)

	port := ":8080"
	log.Printf("Started app on 127.0.0.1%s", port)
	http.ListenAndServe(port, configured)
}

func Connect() *gorm.DB {
	log.Printf("Connecting to Database...")

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
	log.Printf("Migrating Database...")

	db.AutoMigrate(&Feed{})
	db.AutoMigrate(&filestore.Task{})
}

func ToString(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Cannot convert string to int!")
	}
	return i
}

type Response struct {
	Message string `json:"Message"`
}
