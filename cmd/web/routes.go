package main

import (
	"net/http"

	"github.com/antonrodin/trevor-boilerplate/pkg/config"
	"github.com/antonrodin/trevor-boilerplate/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(WriteToTheConsole)
	router.Use(CsrfProitection)
	router.Use(LoadSession)

	// Actual routes
	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	router.Get("/test", handlers.Repo.Test)
	router.Get("/store-ip", handlers.Repo.StoreIp)

	return router
}
