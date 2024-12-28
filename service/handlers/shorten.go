package handlers

import (
	"encoding/json"
	"net/http"
)

type request struct {
	OriginalURL string `json:"originalUrl"`
}

type response struct {
	ShortURL string `json:"shortUrl"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// TODO: Add logic to generate a short alias

	// TODO: Store short alias in Redis

	response := response{
		ShortURL: "https://localhost:8080/redirect/" + "alias123",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
