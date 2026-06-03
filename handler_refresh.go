package main

import (
	"chirpy/internal/auth"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetRefreshToken(r.Header)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid refresh token", err)
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	if refreshToken == "" {
		respondWithError(w, http.StatusUnauthorized, "Missing refresh token", nil)
		return
	}

	rt, err := cfg.queries.GetRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid refresh token", err)
		return
	}

	if rt.ExpiresAt.Before(time.Now()) {
		respondWithError(w, http.StatusUnauthorized, "Refresh token expired", nil)
		return
	}

	jwt, err := auth.MakeJWT(rt.UserID, cfg.jwtSecret)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to make JWT", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{Token: jwt})

}
