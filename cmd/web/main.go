package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antonrodin/trevor-boilerplate/pkg/config"
	"github.com/antonrodin/trevor-boilerplate/pkg/handlers"
	"github.com/antonrodin/trevor-boilerplate/pkg/render"
)

const portNumber = ":3000"

func main() {

	var app config.AppConfig

	templateCache, err := render.CreateTemplateCacheBetter()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = templateCache
	app.UseCache = true

	repo := handlers.CreateNewRepository(&app)
	handlers.SetNewRepo(repo)

	render.SetAppConfig(&app)

	fmt.Println("Server is running at http://localhost:3000")

	server := &http.Server{
		Addr:    ":3000",
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
