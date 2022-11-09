package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	config2 "github.com/prashanth-gajula/bookings/pkg/config"
	handlers2 "github.com/prashanth-gajula/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config2.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers2.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers2.Repo.About))
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers2.Repo.Home)
	mux.Get("/about", handlers2.Repo.About)
	return mux

}
