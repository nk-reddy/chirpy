package endpoints

import (
	"encoding/json"
	"net/http"
	"strings"
)

func HandlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Chirp string `json:"body"`
	}
	type errorVal struct {
		Error string `json:"error"`
	}
	type normVal struct {
		CleanedBody string `json:"cleaned_body"`
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

		text := cleanString(params.Chirp)
		respondWithJSON(w, 200, normVal{CleanedBody: text})
	}
}

func cleanString(str string) string {
	words := strings.Split(str, " ")
	for i := range words {
		word := strings.ToLower(words[i])
		if word == "kerfuffle" || word == "sharbert" || word == "fornax" {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
