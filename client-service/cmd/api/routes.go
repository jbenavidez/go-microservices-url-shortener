package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/url-shortener", a.CreateUser)

	return mux
}
