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
		read_list_items.status
		FROM read_list_items
		JOIN books
			ON books.id = read_list_items.book_id;`

	rows, err := tools.DB.Query(queryString)
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
			&item.Status,
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

func handleAddToReadListSQLError(err error) (error, int) {
	log.Error(err)
	var errMsg string = fmt.Sprint(err)
	var status int = 500
	var returnErr error = errors.New("Sorry, something went wrong adding this book to your read list.")

	if strings.Contains(errMsg, "read_list_items.unique_user_book") {
		returnErr = errors.New("Sorry, it appears this book is already on your read list!")
		status = 400
	} else if strings.Contains(errMsg, "Data truncated for column 'status'") {
		returnErr = errors.New("Sorry, it appears you've provided an invalid status.")
		status = 400
	}

	return returnErr, status
}

func AddBookToReadlist (bookId int, userId int, request resources.AddBookToReadlistRequest) (error, int) {
	var insertQuery string = `
	INSERT INTO read_list_items
	(status, user_id, book_id)
	VALUES
	(?, ?, ?)`

	_, err := tools.DB.Exec(insertQuery, request.Status, userId, bookId)
	
	if err != nil {
		err, code := handleAddToReadListSQLError(err)
		return err, code
	}

	return nil, -1
}
