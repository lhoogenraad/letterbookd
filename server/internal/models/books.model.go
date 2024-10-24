package models

import (
	"database/sql"
	"errors"
	"fmt"
	"server/internal/resources"
	"server/internal/tools"
	"time"
)


func GetBooks(userId int) ([]resources.BookData, error) {
	var queryString string = `
	SELECT 
		books.id,
		books.name, 
		CONCAT(authors.first_name, ' ', authors.last_name) as author_name,
		books.published_date, books.num_pages, IFNULL(books.cover_url, ''),
		IFNULL(books.synopsis, 'No synopsis.'),
		read_list_items.id IS NOT NULL
	FROM books
	JOIN authors
		ON books.author_id=authors.id
	LEFT JOIN read_list_items
		ON read_list_items.book_id = books.id
		AND read_list_items.user_id = ?;
	`

	rows, err := tools.DB.Query(queryString, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []resources.BookData
	books, err = readBookRows(rows)

	if err = rows.Err(); err != nil {
		return books, err
	}
	return books, nil
}


func readBookRows (rows *sql.Rows) ([]resources.BookData, error) {
	var books []resources.BookData

	for rows.Next() {
		var book resources.BookData
		var date string
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&date,
			&book.NumPages,
			&book.CoverURL,
			&book.Synopsis,
			&book.OnUserReadlist,
		); err != nil {
			return books, err
		}
		var err error
		book.Published, err = time.Parse("2006-01-02", date)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	return books, nil
}


func GetSingleBook(bookId int, userId int) (resources.BookData, error, int) {
	var queryString string = `
	SELECT 
		books.id,
		books.name, 
		CONCAT(authors.first_name, ' ', authors.last_name) as author_name,
		books.published_date, 
		books.num_pages, 
		IFNULL(books.cover_url, ''),
		IFNULL(books.synopsis, 'No synopsis.'),
		read_list_items.id IS NOT NULL
	FROM books
	JOIN authors
		ON books.author_id=authors.id
	LEFT JOIN read_list_items
		ON read_list_items.book_id = books.id
		AND read_list_items.user_id = ?
	WHERE books.id = ?;`

	row := tools.DB.QueryRow(queryString, userId, bookId)

	var book resources.BookData
	var date string
	switch err := row.Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&date,
		&book.NumPages,
		&book.CoverURL,
		&book.Synopsis,
		&book.OnUserReadlist,
	); err {
	case sql.ErrNoRows:
		return book, errors.New(`Could not find book`), 404
	}
	fmt.Println(`Retrieved book: `, book.Title)
	var err error
	book.Published, err = time.Parse("2006-01-02", date)
	if err != nil {
		return book, err, 500
	}
	return book, nil, 0
}



func GetBookAverageRating(bookId int, output chan <- float64) (float64, error) {
	var selectQuery string = `
	SELECT IFNULL(AVG(rating), 0)
	FROM reviews
	WHERE book_id=?`

	row := tools.DB.QueryRow(selectQuery, bookId)

	var avgReview float64
	err := row.Scan(&avgReview)
	if err != nil {
		return -1, errors.New(`Something went wrong on our end. Please try again later`)
	}

	output <- float64(avgReview)
	return avgReview, nil
}



func GetBookNumberReviews(bookId int, output chan <- int) (int, error) {
	var selectQuery string = `
	SELECT IFNULL(COUNT(id), 0)
	FROM reviews
	WHERE book_id=?`

	row := tools.DB.QueryRow(selectQuery, bookId)

	var numReviews int
	err := row.Scan(&numReviews)
	if err != nil {
		return -1, errors.New(`Something went wrong on our end. Please try again later`)
	}

	output <- numReviews
	return numReviews, nil
}


func GetBookCompletedReads(bookId int, output chan <- int) (int, error) {
	var selectQuery string = `
	SELECT IFNULL(COUNT(id), 0)
	FROM read_list_items
	WHERE book_id=?
	AND status='Read'`

	row := tools.DB.QueryRow(selectQuery, bookId)

	var numReads int
	err := row.Scan(&numReads)
	if err != nil {
		return -1, errors.New(`Something went wrong on our end. Please try again later`)
	}

	output <- numReads
	return numReads, nil
}


func GetBookReadlistOccurences(bookId int, output chan <- int) (int, error) {
	var selectQuery string = `
	SELECT IFNULL(COUNT(id), 0)
	FROM read_list_items
	WHERE book_id=?`

	row := tools.DB.QueryRow(selectQuery, bookId)

	var numOccurences int
	err := row.Scan(&numOccurences)
	if err != nil {
		return -1, errors.New(`Something went wrong on our end. Please try again later`)
	}

	output <- numOccurences
	return numOccurences, nil
}
