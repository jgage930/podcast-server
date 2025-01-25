package filestore

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// Download everything that has been pushed into the queue.
// Download everything syncrhonously for now.
func RunQueue(queue *Queue) {
	if queue.IsEmpty() {
		log.Println("Nothing to download, queue is empty!")
		return
	}

	log.Printf("Downloading From Queue...")

	// var failures []Task
	for !queue.IsEmpty() {
		task := queue.Pop()

		// Call url
		r, err := http.Get(task.Url)
		if err != nil {
			log.Printf("Failed to download %s", task.Url)
			continue

			// task.Status = "Failed"
			// failures = append(failures, task)
		}

		// Read body
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("Failed to read body for %s", task.Url)
			continue

			// task.Status = "Failed"
			// failures = append(failures, task)
		}

		// Write to file
		file_store_path := "files/"

		file_name := file_store_path + uuid.New().String() + ".mp3"
		file, err := os.Create(file_name)
		if err != nil {
			log.Printf("Failed to create file for %s", task.Url)
			continue

			// task.Status = "Failed"
			// failures = append(failures, task)
		}

		n, err := file.Write(body)
		if err != nil {
			log.Printf("Failed to write to file for %s", task.Url)
			continue

			// task.Status = "Failed"
			// failures = append(failures, task)
		}

		log.Printf("Wrote %d bytes \n", n)

		file.Sync()
	}

}
