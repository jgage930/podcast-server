package main

import (
	"log"
	"net/http"

	"github.com/mmcdole/gofeed"
)

type Podcast struct{}

func parseIntoPodcast(url string) Podcast {
	parser := gofeed.NewParser()
	feed, _ := parser.ParseURL(url)

	log.Println(feed.Title)

	return Podcast{}
}

func PodcastRouter(h *PodcastHandler, mux *http.ServeMux) {
	mux.HandleFunc("POST /parse", h.parseFeed)
}

type PodcastHandler struct{}

type ParseParameters struct {
	FeedId string `json:"feed_id"`
}

func (*PodcastHandler) parseFeed(w http.ResponseWriter, r *http.Request) {
	var payload ParseParameters
	ReadBody(&payload, w, r)

	log.Print(payload.FeedId)
	log.Printf("called for id %s", payload.FeedId)
}
