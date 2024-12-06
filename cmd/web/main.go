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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Server is running at http://localhost:3000")

	_ = http.ListenAndServe(portNumber, nil)
}
