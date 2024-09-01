package models

import (
	"server/internal/tools"
	"server/internal/resources"
	"errors"
	log "github.com/sirupsen/logrus"
)


func GetReviewComments (reviewId int) ([]resources.ReviewComment, error) {
	var selectQuery string = `
	SELECT 
	review_comments.id,
	comment,
	user_id,
	CONCAT(users.first_name, ' ', users.last_name),
	review_id
	FROM review_comments
	JOIN users ON users.id = user_id
	WHERE review_id = ?;`

	rows, err := tools.DB.Query(selectQuery, reviewId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer rows.Close()

	var comments []resources.ReviewComment
	for rows.Next(){
		var comment resources.ReviewComment
		err := rows.Scan(
			&comment.Id,
			&comment.Comment,
			&comment.UserId,
			&comment.Username,
			&comment.ReviewId,
		)

		if err != nil {
			log.Error("Error scanning comment query results:")
			log.Error(err)
			return nil, err
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