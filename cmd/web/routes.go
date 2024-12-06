package main

import (
	"net/http"

	"github.com/antonrodin/trevor-boilerplate/pkg/config"
	"github.com/antonrodin/trevor-boilerplate/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)

	return router
}
