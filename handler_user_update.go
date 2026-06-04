package main

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerUpdatePassword(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	accessToken, err := auth.GetBearerToken(r.Header)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, err := auth.ValidateJWT(accessToken, cfg.jwtSecret)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err = decoder.Decode(&params)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	newPasswordHash, err := auth.HashPassword(params.Password)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user, err := cfg.queries.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:             userId,
		HashedPassword: newPasswordHash,
		Email:          params.Email,
	})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed.Bool,
	})

}
