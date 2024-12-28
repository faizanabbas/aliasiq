package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Analytics(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")

	w.Write([]byte("Analytics for alias: " + alias))
}
