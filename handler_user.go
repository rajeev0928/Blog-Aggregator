package main

import (
	"encoding/json"
	//"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	//"github.com/rajeev0928/GoTest/internal/auth"
	"github.com/rajeev0928/GoTest/internal/database"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	
	decoder := json.NewDecoder(r.Body)
	
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	// apiKey, err := auth.GetAPIKey(r.Header)
	// log.Println(r.Header)
	
	// if err != nil {
	// 	respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
	// 	return
	// }

	// user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	// if err != nil {
	// 	respondWithError(w, http.StatusNotFound, "Couldn't get user")
	// 	return
	// }

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}