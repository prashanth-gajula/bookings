package main

import (
	nosurf2 "github.com/justinas/nosurf"
	"net/http"
)

// NoSurf will provide csrf request to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf2.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad will saves and loads the request on every session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
