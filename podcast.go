package main

import (
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
	"log"
	"net/http"
	"podcast-server/common"
)

type Podcast struct {
	Title string `json:"title"`
}

func parseIntoPodcast(url string) gofeed.Feed {
	parser := gofeed.NewParser()
	feed, _ := parser.ParseURL(url)

	log.Println(feed.Title)

	return *feed
}

func PodcastRouter(h *PodcastHandler, mux *http.ServeMux) {
	mux.HandleFunc("POST /parse", h.parseFeed)
}

type PodcastHandler struct {
	db *gorm.DB
}

type ParseParameters struct {
	FeedId string `json:"feed_id"`
}

func (h *PodcastHandler) parseFeed(w http.ResponseWriter, r *http.Request) {
	var payload ParseParameters
	common.ReadBody(&payload, w, r)

	var feed Feed
	common.GetById(&feed, payload.FeedId, h.db, w)

	parsedFeed := parseIntoPodcast(feed.Url)
	common.Respond(parsedFeed, w)
}
