package main

import (
	"chirpy/internal/database"
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getChirpsHandler(w http.ResponseWriter, r *http.Request) {

	authorId := r.URL.Query().Get("author_id")
	sortOrder := r.URL.Query().Get("sort")

	var chirps []database.Chirp
	var err error

	if authorId != "" {
		chirps, err = cfg.queries.GetChirpsByAuthor(r.Context(), uuid.MustParse(authorId))
	} else {
		chirps, err = cfg.queries.GetChirps(r.Context())
	}

	sort.Slice(chirps, func(i, j int) bool {
		if sortOrder == "desc" {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		}
		return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	parsedChirps := make([]Chirp, len(chirps))

	for i, chirp := range chirps {
		parsedChirps[i] = Chirp{
			ID:        chirp.ID,
			UserID:    chirp.UserID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
		}
	}

	respondWithJSON(w, http.StatusOK, parsedChirps)
}

func (cfg *apiConfig) getChirpHandler(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("id")

	chirp, err := cfg.queries.GetChirp(r.Context(), uuid.MustParse(chirpID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		UserID:    chirp.UserID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
	})
}
