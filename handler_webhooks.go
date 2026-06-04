package main

import (
	"chirpy/internal/auth"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleWebhooks(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	apiKey := auth.GetAPIKey(r.Header)

	if apiKey != cfg.polkaKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithError(w, http.StatusNoContent, "Invalid event", nil)
		return
	}

	userID := params.Data.UserID

	_, err = cfg.queries.UpdateUserStatus(r.Context(), uuid.MustParse(userID))

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, User{})

}
