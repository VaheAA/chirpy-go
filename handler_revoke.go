package main

import (
	"chirpy/internal/auth"
	"net/http"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetRefreshToken(r.Header)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	if token == "" {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	err = cfg.queries.RevokeRefreshToken(r.Context(), token)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
