package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/sindrishtepani/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	// mux is a pointer to a Mux type, but also implements an http.Handler as well
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Printf("type is not http.Handler, but is %T", v))
	}
}
