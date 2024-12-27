package main

import (
	"encoding/json"
	//"log"
	"net/http"
	"time"

	"github.com/google/uuid"
//	"github.com/rajeev0928/GoTest/internal/auth"
	"github.com/rajeev0928/GoTest/internal/database"
)

func (cfg *apiConfig) handlerFeedsCreate(w http.ResponseWriter, r *http.Request,user database.User) {
	type parametersfeeds struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	
	decoderfeed := json.NewDecoder(r.Body)
	
	paramsfeed := parametersfeeds{}

	err :=decoderfeed.Decode(&paramsfeed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters feeds")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      paramsfeed.Name,
		Url   : paramsfeed.URL,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databasFeedToFeed(feed))
}
func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {


	feeds, err := cfg.DB.GetFeed(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get user")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}