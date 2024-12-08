package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToTheConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit page", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func CsrfProitection(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
