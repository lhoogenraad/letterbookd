package models

import (
	"server/internal/tools"
	"database/sql"
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
	} else if strings.Contains(fmt.Sprint(err), "") {
		returnErr = nil
		status = 0
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



func UpdateReview(userId int, reviewId int, req resources.UpdateReviewBody) (error, int) {
	var insertQuery string = `
	UPDATE reviews
	SET 
		rating = ?,
		description = ?
	WHERE
		id = ? AND
		user_id = ?;`

	_, err := tools.DB.Exec(
		insertQuery,
		req.Rating,
		req.Description,
		reviewId,
		userId,
	)

	if err != nil {
		err, code := handleCreateReviewErr(err)
		return err, code
	}

	return nil, -1
}

var addLikeQuery string = `
	INSERT INTO review_likes (user_id, review_id)
	VALUES (?, ?);
`

var removeLikeQuery string = `
	DELETE FROM review_likes 
	WHERE user_id=? AND review_id=?;
`

func SetReviewLike(reviewId int, userId int, on bool) (error, int) {
	var query string
	if on {
		query = addLikeQuery
	} else {
		query = removeLikeQuery
	}

	_, err := tools.DB.Exec(
		query,
		reviewId,
		userId,
	)

	if err != nil {
		err, code := handleCreateReviewErr(err)
		return err, code
	}

	return nil, -1
}


func GetBookReviews(bookId int, userId int) ( []resources.ReviewData, error ) {
	var selectQueryString string = `
	SELECT 
	reviews.id as review_id,
	users.id as user_id,
	CONCAT( users.first_name, ' ', users.last_name) as user_name,
	reviews.description,
	reviews.rating,
	COUNT(DISTINCT(review_comments.id)) as num_comments,
	COUNT(DISTINCT(review_likes.id)) as num_likes,
    MAX(CASE WHEN review_likes.user_id = ? THEN 1 ELSE 0 END) AS has_user_liked,
	reviews.book_id,
	books.name
	FROM reviews
	JOIN users
		ON users.id = reviews.user_id
	LEFT JOIN review_comments
		ON review_comments.review_id=reviews.id
		AND review_comments.archived = false
	LEFT JOIN review_likes
		ON review_likes.review_id=reviews.id
	LEFT JOIN books
		ON reviews.book_id = books.id
	WHERE reviews.book_id = ?

	GROUP BY 
		reviews.id,
		users.id,
		user_name,
		reviews.description,
		reviews.rating;
	`

	rows, err := tools.DB.Query(selectQueryString, userId, bookId)


	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer rows.Close()

	reviews, err := readReviewRows(rows)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return reviews, nil
}

func GetPopularReviews (userId int) ( []resources.ReviewData, error ){
	var selectQueryString string = `
	SELECT 
	reviews.id as review_id,
	users.id as user_id,
	CONCAT( users.first_name, ' ', users.last_name) as user_name,
	reviews.description,
	reviews.rating,
	COUNT(DISTINCT(review_comments.id)) as num_comments,
	IFNULL(COUNT(DISTINCT(review_likes.id)), 0) as num_likes,
    MAX(CASE WHEN review_likes.user_id = ? THEN 1 ELSE 0 END) AS has_user_liked,
	reviews.book_id,
	books.name
	FROM reviews
	JOIN users
		ON users.id = reviews.user_id
	LEFT JOIN review_comments
		ON review_comments.review_id=reviews.id
		AND review_comments.archived = false
	LEFT JOIN review_likes
		ON review_likes.review_id=reviews.id
	LEFT JOIN books
		ON reviews.book_id = books.id

	WHERE review_likes.timestamp BETWEEN (NOW() - INTERVAL 2 WEEK) AND NOW()

	GROUP BY 
		reviews.id,
		users.id,
		user_name,
		reviews.description,
		reviews.rating
	
	ORDER BY num_likes DESC
	LIMIT 10;`

	rows, err := tools.DB.Query(selectQueryString, userId)


	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer rows.Close()

	reviews, err := readReviewRows(rows)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return reviews, nil

}


func checkBookExists (bookId int) bool {
	var getBookQuery string = `
		SELECT * FROM books
		WHERE
			id = ?;`
	
	err := tools.DB.QueryRow(getBookQuery, bookId)

	if err != nil {
		fmt.Println("no book found with that id", err)
		return false
	}

	return true
}


func readReviewRows (rows *sql.Rows) ([]resources.ReviewData, error) {
	var reviews []resources.ReviewData
	for rows.Next() {
		var review resources.ReviewData
		err := rows.Scan(
			&review.Id,
			&review.UserId,
			&review.Username,
			&review.Description,
			&review.Rating,
			&review.NumComments,
			&review.NumLikes,
			&review.LikedBy,
			&review.BookId,
			&review.BookTitle,
		)

		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
