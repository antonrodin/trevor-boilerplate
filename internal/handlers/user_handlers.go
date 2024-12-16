package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/antonrodin/trevor-boilerplate/internal/forms"
	"github.com/antonrodin/trevor-boilerplate/internal/models"
	"github.com/antonrodin/trevor-boilerplate/internal/render"
	"github.com/justinas/nosurf"
)

func (m *Repository) UserCreate(w http.ResponseWriter, r *http.Request) {

	type TemplateData struct {
		CsrfToken string
		Form      *forms.Form
	}

	data := TemplateData{
		nosurf.Token(r),
		forms.New(nil),
	}

	err := render.GetTemplate("user/create.page.go.tmpl").Execute(w, data)

	if err != nil {
		log.Fatal(err)
	}
}

func (m *Repository) UserStore(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panicln("Error on parse user form", err)
		return
	}

	user := models.User{
		Username: r.Form.Get("username"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	//form := forms.New(r.PostForm)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
