package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Arsh910/go/internal/database"
	"github.com/google/uuid"
)

func (apic *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string
	}

	decoder := json.NewDecoder(r.Body)

	params := Params{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Errro: %s", err))
		return
	}

	user, err := apic.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error: %s", err))
		return
	}
	respondWithJson(w, 200, databaseUserToUser(user))
}
