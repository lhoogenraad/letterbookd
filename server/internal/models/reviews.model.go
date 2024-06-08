package models

import (
	"server/internal/tools"
	"server/internal/resources"
	"strings"
	"fmt"
	"errors"
	log "github.com/sirupsen/logrus"
)


func handleCreateReviewErr(err error) (error, int) {
	log.Error(err)
	returnErr := errors.New("Sorry, something went wrong while submitting your review.")
	status := 500

	if strings.Contains(fmt.Sprint(err), "reviews_ibfk_1") {
		returnErr = errors.New("Sorry, it seems we can't create a review for this book.")
		status = 400
	} else if strings.Contains(fmt.Sprint(err), "user_book_review_unique") {
		returnErr = errors.New("Sorry, it seems you already have made a review for this book. You may only leave one review per book.")
		status = 400
	}
	return returnErr, status
}


func CreateReview(userId int, bookId int, req resources.CreateReviewBody) (error, int) {
	var insertQuery string = `
	INSERT into reviews
	(user_id, book_id, rating, description)
	VALUES
	(?, ?, ?, ?)`

	_, err := tools.DB.Exec(
		insertQuery,
		userId,
		bookId,
		req.Rating,
		req.Description,
	)

	if err != nil {
		err, code := handleCreateReviewErr(err)
		return err, code
	}

	return nil, -1
}
