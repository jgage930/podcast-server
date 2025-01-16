package main

import (
  "net/http"
  "log"
)

func main() {
  port := ":8080"
  log.Printf("Starting Application.")

  http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Podcast Server v0.0.1"))
  })

  err := http.ListenAndServe(port, nil)
  if err != nil {
    log.Fatal(err)
  }
}
