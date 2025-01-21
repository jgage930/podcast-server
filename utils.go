package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Payload interface{}

func ReadBody[T Payload](buf *T, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Failed to read request body.", http.StatusBadRequest)
		return
	}

	log.Print(body)

	if err := json.Unmarshal(body, &buf); err != nil {
		http.Error(w, "Invalid Json Body", http.StatusUnprocessableEntity)
	}
}

type JsonResponse interface{}

func Respond[T JsonResponse](body T, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}
