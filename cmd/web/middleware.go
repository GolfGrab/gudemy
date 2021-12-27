package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is a middleware that protects against Cross-Site Request Forgery
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoader is a middleware that loads the session from the request
func SessionLoader(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
