package models

import (
	"github.com/lhoogenraad/letterbookd/internal/tools"
	"github.com/lhoogenraad/letterbookd/internal/resources"
	"errors"
	"time"
	"fmt"
	log "github.com/sirupsen/logrus"
)


func GetReviewComments (reviewId int) ([]resources.ReviewComment, error) {
	var selectQuery string = `
	SELECT 
	review_comments.id,
	comment,
	user_id,
	CONCAT(users.first_name, ' ', users.last_name),
	review_id,
	edited,
	timestamp
	FROM review_comments
	JOIN users ON users.id = user_id
	WHERE review_id = ?
	AND archived = false;`

	rows, err := tools.DB.Query(selectQuery, reviewId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer rows.Close()

	var comments []resources.ReviewComment
	for rows.Next(){
		var comment resources.ReviewComment
		var date string
		err := rows.Scan(
			&comment.Id,
			&comment.Comment,
			&comment.UserId,
			&comment.Username,
			&comment.ReviewId,
			&comment.Edited,
			&date,
		)

		if err != nil {
			log.Error("Error scanning comment query results:")
			log.Error(err)
			return nil, err
		}
		fmt.Print(date)
		comment.Timestamp, err = time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			return comments, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}


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

func DeleteReviewComment (reviewId int, userId int, commentId int) (error, int) {
	var deleteQuery = `
	UPDATE review_comments
	SET archived=true
	WHERE
		id = ? AND
		user_id = ? AND
		review_id = ?;
	`

	_, err := tools.DB.Exec(deleteQuery, commentId, userId, reviewId)
	
	if err != nil {
		log.Error(err)

		return errors.New("Sorry, something went wrong leaving your comment"), 500
	}

	return nil, -1
}


func UpdateReviewComment (reviewId int, userId int, commentId int, request resources.CreateReviewCommentBody) (error, int) {
	var insertQuery string = `
	UPDATE review_comments
	SET 
		comment=?,
		edited=true
	WHERE
		id=? AND
		user_id=? AND
		review_id=?`

	_, err := tools.DB.Exec(insertQuery, request.Comment, commentId, userId, reviewId)
	
	if err != nil {
		log.Error(err)

		return errors.New("Sorry, something went wrong updating your comment"), 500
	}

	return nil, -1
}
