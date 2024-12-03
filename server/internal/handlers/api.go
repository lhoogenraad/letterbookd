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
		AllowedOrigins:   []string{"http://localhost:3000", "https://deluxe-malabi-dfe5af.netlify.app", "https://deluxe-malabi-dfe5af.netlify.app/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{
			"Accept",
			"Accept-Encoding",
			"Authorization", 
			"Connection",
			"Content-Type", 
			"User-Agent",
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

	r.Route("/api", func (router chi.Router) {
		router.Get("/healthcheck", HelloJamie)
	})

	r.Route("/api", func(router chi.Router) {

		// Middleware for auth on these sensitive routes
		router.Use(middleware.Authorisation)
		router.Use(middleware.Paginate)

		// Books
		router.Get("/books", GetBooks)
		router.Get("/books/count", GetBooksCount)
		router.Get("/books/featured", GetFeaturedBooks)
		router.Get("/books/search/open-library", SearchOpenLibrary)
		router.Post("/books/search/open-library/{olId}", ConfirmOpenLibraryBookUpload)
		router.Get("/books/{bookId}", GetSingleBook)

		// Readlist
		router.Get("/read-list", GetReadListItems)
		router.Post("/read-list/book/{bookId}", AddBookToReadList)
		router.Put("/read-list/book/{bookId}", UpdateReadListItem)
		router.Delete("/read-list/book/{bookId}", DeleteReadListItem)

		// Reviews
		router.Post("/books/{bookId}/reviews", CreateReview)
		router.Get("/reviews/popular", GetPopularReviews)
		router.Put("/reviews/{reviewId}", UpdateReview)
		router.Get("/books/{bookId}/reviews", GetBookReviews)
		router.Get("/books/{bookId}/reviews/summary", GetBookReviewSummary)

		// Review Likes
		router.Post("/reviews/{reviewId}/like", AddReviewLike)
		router.Post("/reviews/{reviewId}/unlike", RemoveReviewLike)

		// Review Comments
		router.Post("/reviews/{reviewId}/comments", CreateReviewComment)
		router.Get("/reviews/{reviewId}/comments", GetReviewComments)
		router.Delete("/reviews/{reviewId}/comments/{commentId}", DeleteReviewComment)
		router.Put("/reviews/{reviewId}/comments/{commentId}", UpdateReviewComment)

	})
}
