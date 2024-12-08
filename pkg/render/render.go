package render

import (
	"html/template"
	"log"
	"net/http"

	"github.com/antonrodin/trevor-boilerplate/pkg/config"
)

var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var parsedTemplate *template.Template
	var err error

	parsedTemplate, err = template.ParseFiles("./templates/"+tmpl, "./templates/main.layout.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTemplate(tmpl string) *template.Template {
	var parsedTemplate *template.Template
	var err error

	parsedTemplate, err = template.ParseFiles("./templates/"+tmpl, "./templates/main.layout.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	return parsedTemplate
}
