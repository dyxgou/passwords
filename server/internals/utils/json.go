package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithErr(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Printf("Failed with %d code, msg : %s\n", code, msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON %v\n", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
