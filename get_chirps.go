package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/nk-reddy/chirpy/internal/database"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	chirps := []database.Chirp{}
	var err error
	s := r.URL.Query().Get("author_id")
	if s != "" {
		authorID, err := uuid.Parse(s)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "invalid author ID")
			return
		}
		chirps, err = cfg.db.GetChirpsByAuthor(r.Context(), authorID)
	} else {
		chirps, err = cfg.db.GetChirps(r.Context())
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := make([]chirpResponse, len(chirps))

	for i, chirp := range chirps {
		response[i] = chirpResponse{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
	}

	s2 := r.URL.Query().Get("sort")
	if s2 == "desc" {
		sort.Slice(response, func(i, j int) bool {
			return response[i].CreatedAt.After(response[j].CreatedAt)
		})
	} else {
		sort.Slice(response, func(i, j int) bool {
			return response[i].CreatedAt.Before(response[j].CreatedAt)
		})
	}
	respondWithJSON(w, http.StatusOK, response)
}
