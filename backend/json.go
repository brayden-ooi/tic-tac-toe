package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to marshal payload to JSON: ", err)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, status int, err string) {
	if status > 499 {
		log.Println("Responding with 5XX error: ", err)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, status, errResponse{
		Error: err,
	})
}
