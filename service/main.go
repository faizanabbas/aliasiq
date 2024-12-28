package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, from AliasIQ."))
	})

	r.Post("/shorten", shortenURLHandler)
	r.Get("/redirect/{alias}", redirectURLHandler)
	r.Get("/analytics/{alias}/", analyticsHandler)

	http.ListenAndServe(":8080", r)
}
func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shorten URL"))
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Redirect URL"))
}

func analyticsHandler(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")

	w.Write([]byte("Analytics for alias: " + alias))
}
