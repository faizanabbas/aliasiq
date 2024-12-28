package main

import (
	"github.com/faizanabbas/aliasiq/service/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(r chi.Router) {
	r.Post("/shorten", handlers.ShortenURL)
	r.Get("/redirect/{alias}", handlers.RedirectURL)
	r.Get("/analytics/{alias}", handlers.Analytics)
}
