package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/antonrodin/trevor-boilerplate/pkg/config"
	"github.com/antonrodin/trevor-boilerplate/pkg/handlers"
	"github.com/antonrodin/trevor-boilerplate/pkg/render"
)

const portNumber = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Some app config
	app.InProduction = false
	app.AppTitle = "Trevor Boilerplate"

	// Session package
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// Persist if browser is closed
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// Check https or http
	session.Cookie.Secure = app.InProduction

	// Add session to app
	app.Session = session

	// New repository so we can send global configuration to handlers
	repo := handlers.CreateNewRepository(&app)
	handlers.SetNewRepo(repo)
	render.SetAppConfig(&app)

	fmt.Println("Server is running at http://localhost:3000")

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
