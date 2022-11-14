package main

import (
	"fmt"
	"github.com/go-chi/chi"
	config2 "github.com/prashanth-gajula/bookings/internal/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config2.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Error(fmt.Sprintf("Type is not *chi.mux type is: %T", v))

	}

}
