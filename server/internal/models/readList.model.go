package models

import (
	"server/internal/resources"
	"server/internal/tools"
	"errors"
	"strings"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func GetReadListItems(userId int) ([]resources.ReadListItem, error) {
	var queryString string = `
		SELECT
		read_list_items.id,
		read_list_items.user_id,
		books.id,
		books.name,
		IFNULL(books.synopsis, ''),
		read_list_items.status,
		reviews.id IS NOT NULL
		FROM read_list_items
		JOIN books
			ON books.id = read_list_items.book_id
		LEFT JOIN reviews
			ON reviews.book_id = books.id
			AND reviews.user_id = ?
		WHERE read_list_items.user_id = ?
		;`

	rows, err := tools.DB.Query(queryString, userId, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var readListItems []resources.ReadListItem

	for rows.Next() {
		var item resources.ReadListItem
		if err := rows.Scan(
			&item.Id,
			&item.UserId,
			&item.BookId,
			&item.BookName,
			&item.BookSynopsis,
			&item.Status,
			&item.UserReviewed,
		); err != nil {
			return readListItems, err
		}
		readListItems = append(readListItems, item)
	}
	if err = rows.Err(); err != nil {
		return readListItems, err
	}
	return readListItems, nil
}

func handleReadListModSQLError(err error) (error, int) {
	log.Error(err)
	var errMsg string = fmt.Sprint(err)
	var status int = 500
	var returnErr error = errors.New("Sorry, something went wrong.")

	if strings.Contains(errMsg, "read_list_items.unique_user_book") {
		returnErr = errors.New("Sorry, it appears this book is already on your read list!")
		status = 400
	} else if strings.Contains(errMsg, "Data truncated for column 'status'") {
		returnErr = errors.New("Sorry, it appears you've provided an invalid status.")
		status = 400
	}

	return returnErr, status
}

func AddBookToReadlist (bookId int, userId int, request resources.ReadListModReq) (error, int) {
	var insertQuery string = `
	INSERT INTO read_list_items
	(status, user_id, book_id)
	VALUES
	(?, ?, ?)`

	_, err := tools.DB.Exec(insertQuery, request.Status, userId, bookId)
	
	if err != nil {
		err, code := handleReadListModSQLError(err)
		return err, code
	}

	return nil, -1
}


func UpdateReadListItem (bookId int, userId int, request resources.ReadListModReq) (error, int) {
	var updateQuery string = `
	UPDATE read_list_items
	SET status = ?
	WHERE user_id = ?
	AND book_id = ?`

	_, err := tools.DB.Query(updateQuery, request.Status, userId, bookId)
	
	if err != nil {
		err, code := handleReadListModSQLError(err)
		return err, code
	}

	return nil, -1
}



func DeleteReadListItem (bookId int, userId int) (error, int) {
	var updateQuery string = `
	DELETE FROM read_list_items
	WHERE user_id = ?
	AND book_id = ?;`

	_, err := tools.DB.Query(updateQuery, userId, bookId)
	
	if err != nil {
		err, code := handleReadListModSQLError(err)
		return err, code
	}

	return nil, -1
}
