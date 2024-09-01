package models

import (
	"server/internal/tools"
	"server/internal/resources"
	"errors"
	log "github.com/sirupsen/logrus"
)


func CreateReviewComment (reviewId int, userId int, request resources.CreateReviewCommentBody) (error, int) {
	var insertQuery string = `
	INSERT INTO review_comments
	(comment, user_id, review_id)
	VALUES
	(?, ?, ?)`

	_, err := tools.DB.Exec(insertQuery, request.Comment, userId, reviewId)
	
	if err != nil {
		log.Error(err)

		return errors.New("Sorry, something went wrong leaving this review"), 500
	}

	return nil, -1
}
