package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/api"

	// "time"
	"server/internal/models"
	"server/internal/resources"
	"server/internal/utils"
	"strconv"

	log "github.com/sirupsen/logrus"
)


func GetBooks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []resources.BookData
	var err error
	books, err = models.GetBooks()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	err = json.NewEncoder(w).Encode(books)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func GetSingleBook (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookIdParam := utils.GetParam(r, "bookId")
	bookId, err := strconv.Atoi(bookIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid book ID was given as a parameter.")
		return
	}

	var books resources.BookData
	books, err, errCode := models.GetSingleBook(bookId)

	if err != nil {
		log.Error(err)
		api.CustomErrorHandler(w, errCode, fmt.Sprint(err))
		return
	}

	err = json.NewEncoder(w).Encode(books)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}


func GetBookReviewSummary (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookIdParam := utils.GetParam(r, "bookId")
	bookId, err := strconv.Atoi(bookIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid book ID was given as a parameter.")
		return
	}

	avgRating, err := models.GetBookNumberReviews(bookId)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	err = json.NewEncoder(w).Encode(avgRating)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
