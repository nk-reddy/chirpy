package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Chirp string `json:"body"`
	}
	type errorVal struct {
		Error string `json:"error"`
	}
	type normVal struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if len(params.Chirp) > 140 {
		respondWithJSON(w, 400, errorVal{Error: "Chirp is too long"})
	} else {
		respondWithJSON(w, 200, normVal{Valid: true})
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	log.Printf("Error decoding parameter: %s", msg)
	w.WriteHeader(code)
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	dat, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}
