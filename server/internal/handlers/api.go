package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"server/internal/middleware"
)

func ApiHandler(r *chi.Mux) {
	// middlewhere to strip trailing slashes
	r.Use(chimiddle.StripSlashes)

	r.Route("/api/users", func(router chi.Router) {
		router.Post("/signup", Signup)
	})

	r.Route("/api", func(router chi.Router) {

		// Middleware for auth on these sensitive routes
		router.Use(middleware.Authorisation)

		router.Get("/books", GetBooks)
	})
}
