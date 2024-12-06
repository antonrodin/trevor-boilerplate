package handlers

import (
	"net/http"

	"github.com/antonrodin/trevor-boilerplate/pkg/config"
	"github.com/antonrodin/trevor-boilerplate/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func CreateNewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func SetNewRepo(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.go.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.go.tmpl")
}
