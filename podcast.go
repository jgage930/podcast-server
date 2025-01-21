package main

import (
	"log"

	"github.com/mmcdole/gofeed"
)

type Podcast struct{}

func ParseIntoPodcast(url string) Podcast {
	parser := gofeed.NewParser()
	feed, _ := parser.ParseURL(url)

	log.Println(feed.Title)

	return Podcast{}
}
