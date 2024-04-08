package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"server/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middlewhere
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorisation)

		router.Get("/coins", GetCoinBalance)
	})
}
