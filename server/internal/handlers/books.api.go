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
	// Grab user Id from claims
	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}
	userId := int(claims["userid"].(float64))
	pagination := utils.GetPagination(r)
	filter := utils.GetUrlQuery(r, "searchText")

	var books []resources.BookData
	var err error
	if filter == "" {
		books, err = models.GetBooks(userId, pagination.Page, pagination.PageSize)
	} else {
		books, err = models.GetBooksWithFilter(userId, pagination.Page, pagination.PageSize, filter)
	}

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


func GetBooksCount (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var booksCount int
	var err error
	filter := utils.GetUrlQuery(r, "searchText")
	if filter == "" {
		booksCount, err = models.GetBooksCount()
	} else {
		booksCount, err = models.GetBooksCountWithFilter(filter)
	}

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	err = json.NewEncoder(w).Encode(booksCount)

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

	// Grab user Id from claims
	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}
	userId := int(claims["userid"].(float64))

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid book ID was given as a parameter.")
		return
	}

	var books resources.BookData
	books, err, errCode := models.GetSingleBook(bookId, userId)

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


func SearchOpenLibrary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book resources.BookDataOL
	book, err := utils.SearchOpenLibrary("20th century american drama")

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	err = json.NewEncoder(w).Encode(book)
}


func ConfirmOpenLibraryBookUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	olId := utils.GetParam(r, "olId")

	var book resources.BookDataOL
	book, err := utils.UploadBookFromOpenLibrary(olId)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	err = json.NewEncoder(w).Encode("Successfully uploaded " + book.Title)
}


func GetBookReviewSummary (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookIdParam := utils.GetParam(r, "bookId")
	bookId, err := strconv.Atoi(bookIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid book ID was given as a parameter.")
		return
	}

	var summary map[string]interface{} = getBookStats(bookId)

	err = json.NewEncoder(w).Encode(summary)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func getBookStats(bookId int) map[string]interface{} {
	avgRatingChan := make(chan float64, 1)
	numReviewsChan := make(chan int, 1)
	numCompletedReads := make(chan int, 1)
	numReadlistOccurences := make(chan int, 1)

	go models.GetBookAverageRating(bookId, avgRatingChan)
	go models.GetBookNumberReviews(bookId, numReviewsChan)
	go models.GetBookCompletedReads(bookId, numCompletedReads)
	go models.GetBookReadlistOccurences(bookId, numReadlistOccurences)


	var avgRating float64 = <- avgRatingChan
	var numReviews int = <- numReviewsChan
	var completedReads int = <- numCompletedReads
	var readlistOccurences int = <- numReadlistOccurences

	summary := map[string]interface{}{
		"avgRating":avgRating, 
		"numReviews":numReviews,
		"numCompletedReads":completedReads,
		"numReadlistOccurences":readlistOccurences,
	}

	return summary
}
