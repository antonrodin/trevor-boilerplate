// Handlers
package handlers

import (
	"fmt"
	"log"
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
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}

	err := render.GetTemplate("home.page.go.tmpl").Execute(w, sweaters)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	type About struct {
		Title string
	}

	data := About{"About Page"}

	err := render.GetTemplate("about.page.go.tmpl").Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Repository) Test(w http.ResponseWriter, r *http.Request) {
	type TemplateData struct {
		Material string
		RemoteIp string
		Count    uint
	}
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	data := TemplateData{"wool", remoteIp, 17}

	err := render.GetTemplate("test.page.go.tmpl").Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Repository) StoreIp(w http.ResponseWriter, r *http.Request) {

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	fmt.Println("Context", r.Context())

	err := render.GetTemplate("store-ip.page.go.tmpl").Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
