package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"server/internal/middleware"
	cors "github.com/go-chi/cors"
)

func ApiHandler(r *chi.Mux) {
	// middlewhere to strip trailing slashes
	r.Use(chimiddle.StripSlashes)

	// CORS setup
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"http://localhost:3000"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api/users", func(router chi.Router) {
		router.Post("/signup", Signup)
		router.Post("/login", Signin)
	})

	r.Route("/api", func(router chi.Router) {

		// Middleware for auth on these sensitive routes
		router.Use(middleware.Authorisation)

		// Books
		router.Get("/books", GetBooks)

		// Reviews
		router.Post("/books/{bookId}/reviews", CreateReview)
		router.Put("/reviews/{reviewId}", UpdateReview)
		router.Get("/books/{bookId}/reviews", GetBookReviews)

		// Review Comments
		router.Post("/reviews/{reviewId}/comments", CreateReviewComment)
		router.Get("/reviews/{reviewId}/comments", GetReviewComments)
	})
}
