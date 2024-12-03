package handlers

import (
    "github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
    cors "github.com/go-chi/cors"
    "net/http"
    log "github.com/sirupsen/logrus"
    "server/internal/middleware"
)

func ApiHandler(r *chi.Mux) {
    // Middleware to log requests for debugging
    r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            log.Printf("Origin: %s, Method: %s, Headers: %s", r.Header.Get("Origin"), r.Method, r.Header.Get("Access-Control-Request-Headers"))
            next.ServeHTTP(w, r)
        })
    })

    // CORS middleware (should be applied early)
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{
            "http://localhost:3000",
            "https://deluxe-malabi-dfe5af.netlify.app",
        },
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{
            "Accept", "Authorization", "Content-Type", "timeout", "X-CSRF-Token",
        },
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           86400, // Cache preflight for 24 hours
    }))

    // Strip trailing slashes middleware
    r.Use(chimiddle.StripSlashes)

    // Handle preflight requests globally
    r.Options("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")
        w.WriteHeader(http.StatusNoContent)
    })

    // Routes
    r.Route("/api", func(router chi.Router) {
        router.Get("/healthcheck", HelloJamie)

        router.Route("/users", func(router chi.Router) {
            router.Post("/signup", Signup)
            router.Post("/login", Signin)
        })

        router.Route("/", func(router chi.Router) {
            // Middleware for protected routes
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
    })
}

