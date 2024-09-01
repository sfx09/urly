package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, msg any) {
	data, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to marshal response json")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	type Response struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, status, Response{Error: msg})
}
