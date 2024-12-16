package main

import (
	"net/http"

	"github.com/antonrodin/trevor-boilerplate/internal/config"
	"github.com/antonrodin/trevor-boilerplate/internal/handlers"
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
	router.Get("/form", handlers.Repo.Form)
	router.Post("/form", handlers.Repo.Process)

	// User
	router.Get("/user/create", handlers.Repo.UserCreate)
	router.Post("/user/store", handlers.Repo.UserStore)

	return router
}
