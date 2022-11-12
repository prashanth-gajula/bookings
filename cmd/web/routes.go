package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	config2 "github.com/prashanth-gajula/bookings/internal/config"
	handlers2 "github.com/prashanth-gajula/bookings/internal/handlers"
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
	mux.Get("/generals-quarters", handlers2.Repo.Generals)
	mux.Get("/majors-suite", handlers2.Repo.Majors)
	mux.Get("/search-availability", handlers2.Repo.Availability)
	mux.Get("/contact.html", handlers2.Repo.Contact)
	mux.Get("/make-reservation", handlers2.Repo.Reservation)
	mux.Post("/make-reservation", handlers2.Repo.PostReservation)
	mux.Post("/search-availability", handlers2.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers2.Repo.Availabilityjson)
	mux.Get("/Reservation-Summary", handlers2.Repo.ReservationSummary)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
