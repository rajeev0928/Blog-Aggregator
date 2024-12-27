package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type urLToFeed struct {
	Channel struct {
		Title       string          `xml:"title"`
		Link        string          `xml:"link"`
		Description string          `xml:"description"`
		Language    string          `xml:"language"`
		Item        []urLToFeedItem `xml:"item"`
	} `xml:"channel"`
}

type urLToFeedItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(feedURL string) (*urLToFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Get(feedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed urLToFeed
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}