package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/antonrodin/trevor-boilerplate/pkg/config"
)

var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var t *template.Template
	var error error
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCacheBetter()
	}

	t = templateCache[tmpl]

	error = t.Execute(w, nil)

	if error != nil {
		log.Fatal("Error parsing template", error)
		return
	}
}

func CreateTemplateCacheBetter() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.go.tmpl")
	if err != nil {
		return myCache, err
	}

	// Iterate all pages
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob("./templates/*layout.go.tmpl")
		if err != nil {
			log.Println("Error on globing layouts")
			return myCache, nil
		}

		if len(layouts) > 0 {
			tmpl, err = tmpl.ParseGlob("./templates/*layout.go.tmpl")
			if err != nil {
				log.Println("Error on parsing glob layouts")
				return myCache, nil
			}
		}

		myCache[name] = tmpl
	}

	log.Println("Finish generate template cache")
	log.Println(myCache)

	return myCache, nil
}
