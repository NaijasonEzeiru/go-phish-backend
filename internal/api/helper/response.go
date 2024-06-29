package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}

	RespondWithJSON(w, code, ErrResponse{
		Error: msg, StatusCode: code,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
	println(dat)
}
