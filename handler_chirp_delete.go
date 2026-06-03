package main

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) deleteChirpHandler(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")

	type response struct {
		message string
	}

	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	userId, err := auth.ValidateJWT(accessToken, cfg.jwtSecret)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	chirp, err := cfg.queries.GetChirp(r.Context(), uuid.MustParse(chirpID))

	if err != nil {
		http.Error(w, "Chirp not found", http.StatusNotFound)
		return
	}

	if chirp.UserID != userId {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	err = cfg.queries.DeleteChirp(r.Context(), database.DeleteChirpParams{
		ID:     uuid.MustParse(chirpID),
		UserID: userId,
	})

	if err != nil {
		http.Error(w, "Chirp not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusNoContent, response{message: "Chirp deleted"})

}
