package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
)

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
