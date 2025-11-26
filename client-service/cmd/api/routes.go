package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {

	mux := chi.NewRouter()

	mux.Post("/url-shortener", a.CreateUrlShortener)
	mux.Get("/all-url-shortener", a.GetAllUrlShorteners)
	mux.Put("/url-shortener/{id}", a.UpdateUrlShortener)
	return mux
}
