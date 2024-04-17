package models

import (
	"server/internal/tools"
	"server/internal/resources"
	"time"
)


func GetBooks() ([]resources.BookData, error) {
	var queryString string = `
	SELECT 
	books.name, 
	CONCAT(authors.first_name, ' ', authors.last_name) as author_name,
	books.published_date, books.num_pages, IFNULL(books.cover_url, '')
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
			&book.Title,
			&book.Author,
			&date,
			&book.NumPages,
			&book.CoverURL,
		); err != nil {
			return books, err
		}

		book.Published, err = time.Parse("2006-01-01", date)
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
