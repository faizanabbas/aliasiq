package handlers

import (
	"encoding/json"
	gonanoid "github.com/matoous/go-nanoid/v2"
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
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	alias, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz0123456789", 7)
	if err != nil {
		http.Error(w, "error generating short alias", http.StatusInternalServerError)
		return
	}

	// TODO: Store short alias in Redis

	response := response{
		ShortURL: "https://localhost:8080/redirect/" + alias,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}
