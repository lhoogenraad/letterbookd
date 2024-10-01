package models

import (
	"database/sql"
	"server/internal/resources"
	"server/internal/tools"
	"time"
	"errors"
)


func GetBooks() ([]resources.BookData, error) {
	var queryString string = `
	SELECT 
	books.id,
	books.name, 
	CONCAT(authors.first_name, ' ', authors.last_name) as author_name,
	books.published_date, books.num_pages, IFNULL(books.cover_url, ''),
	IFNULL(books.synopsis, 'No synopsis.')
	FROM books
	JOIN authors
	ON books.author_id=authors.id`

	rows, err := tools.DB.Query(queryString)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
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
		); err != nil {
			return books, err
		}

		book.Published, err = time.Parse("2006-01-02", date)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return books, err
	}
	return books, nil
}



func GetSingleBook(bookId int) (resources.BookData, error, int) {
	var queryString string = `
	SELECT 
	books.id,
	books.name, 
	CONCAT(authors.first_name, ' ', authors.last_name) as author_name,
	books.published_date, books.num_pages, IFNULL(books.cover_url, ''),
	IFNULL(books.synopsis, 'No synopsis.')
	FROM books
	JOIN authors
	ON books.author_id=authors.id
	WHERE books.id = ?
	`

	row := tools.DB.QueryRow(queryString, bookId)

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
	); err {
	case sql.ErrNoRows:
		return book, errors.New(`Could not find book`), 404
	}

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
