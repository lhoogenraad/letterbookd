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
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{
			"Accept",
			"Authorization", 
			"Content-Type", 
			"timeout",
			"X-CSRF-Token", 
			"X-Auth-Token",
		},
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
		router.Get("/books/{bookId}", GetSingleBook)

		// Readlist
		router.Get("/read-list", GetReadListItems)
		router.Post("/read-list/book/{bookId}", AddBookToReadList)
		router.Put("/read-list/book/{bookId}", UpdateReadListItem)

		// Reviews
		router.Post("/books/{bookId}/reviews", CreateReview)
		router.Put("/reviews/{reviewId}", UpdateReview)
		router.Get("/books/{bookId}/reviews", GetBookReviews)

		router.Get("/books/{bookId}/reviews/summary", GetBookReviewSummary)

		// Review Comments
		router.Post("/reviews/{reviewId}/comments", CreateReviewComment)
		router.Get("/reviews/{reviewId}/comments", GetReviewComments)
		router.Delete("/reviews/{reviewId}/comments/{commentId}", DeleteReviewComment)
		router.Put("/reviews/{reviewId}/comments/{commentId}", UpdateReviewComment)

	})
}
