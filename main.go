package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting Application.")

	db := Connect()
	SetupDb(db)

	feedMux := feedRouter(&FeedHandler{db})

	appMux := http.NewServeMux()
	appMux.Handle("/feed", feedMux)

	port := ":8080"
	log.Printf("Started app on 127.0.0.1%s", port)
	http.ListenAndServe(port, appMux)
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
}

type Response struct {
	Message string `json:"Message"`
}
