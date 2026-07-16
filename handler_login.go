package main

import (
	"encoding/json"
	"net/http"

	"github.com/nk-reddy/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := userRequest{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil || !match {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	respondWithJSON(w, http.StatusOK, userResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
