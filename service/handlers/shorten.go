package handlers

import "net/http"

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shorten URL"))
}