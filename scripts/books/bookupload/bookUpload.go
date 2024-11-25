package bookupload

import (
	"fmt"
	"time"
	"scripts/util"
	"scripts/authors"
	"scripts/books"
	"scripts/structs"
)

func getAuthorId (book structs.Book, authorMap map[string]int) ( int , bool ) {
	authors := book.Authors
	if authors == nil || len(authors) < 1 {
		return 0, false
	}

	for _, author := range authors {
		id, ok := authorMap[author.Key]
		if ok {
			return id, true
		}
	}

	return 0, false
}

/**
Will upload books which have authors contained in our
database. These author IDs are grabbed from authors package
*/
func ReadAndUpload() error {
	total, totalOked, failed := 0, 0, 0
	authorMap, err := authors.GetAllAuthorIds()
	if err != nil {
		return err
	}
	million := 1000*1000
	offset := million * 3 + 500000
	amount := million * 5

	// Ten milly
	books, err := books.GetAllBooks(amount, offset)

	if err != nil {
		return err
	}
	fmt.Println("len books:", len(books))
	for i := 0; i < len(books); i++{
		book := books[i]
		total++
		authorId, ok := getAuthorId(book, authorMap)
		if ok {
			totalOked++
			err = InsertBook(book, authorId)
			if err != nil {
				fmt.Println(`Couldn't upload book`, book.Title, `:`, err)
				failed++
			}
		} 
		if i % 100000 == 0{
			fmt.Printf("%d out of %d books checked. %d to go!\n", i, len(books), len(books) - i)
		}
	}
	fmt.Printf("\n\nTotal books: %d\tTotal attempted: %d\tUploaded %d\t %d failed", total, totalOked, totalOked-failed, failed)
	return nil
}



func UpdateBooksSynopses() error{
	bookMap, err:= books.GetBookOpenLibIdMap()
	if err != nil {
		return err
	}

	million := 1000*1000
	offset := million * 5
	amount := million * 5

	// Ten milly
	books, err := books.GetAllBooks(amount, offset)

	if err != nil {
		return err
	}

	for i, book := range books {
		bookId, ok := bookMap[book.Key]
		if ok  && book.Description.Value != ""{
			err = updateBookSynopsis(bookId, book.Description.Value)
			if err != nil {
				fmt.Printf("Error updating %s\t", book.Title)
				fmt.Println(err)
			}
		}
		if i % 100000 == 0{
			fmt.Printf("%d out of %d books checked. %d to go!\n", i, len(books), len(books) - i)
		}
	}
	return nil
}




func parseBookDateString(book structs.Book) ( time.Time, error ) {
	var format string
	var date time.Time

	// Commonly the Pub dates are just the year
	if len(book.Publish_date) == 4 {
		format = "2006"
	} else {
		format = "January 02, 2006"
	}

	date, err := time.Parse(format, book.Publish_date)

	if err != nil {
		format = "January 2, 2006"
		date, err = time.Parse(format, book.Publish_date)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "January 2006"
		date, err = time.Parse(format, book.Publish_date)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "Jan 2, 2006"
		date, err = time.Parse(format, book.Publish_date)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "Jan 02, 2006"
		date, err = time.Parse(format, book.Publish_date)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2006-01-02"
		date, err = time.Parse(format, book.Publish_date)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2006-Jan-02"
		date, err = time.Parse(format, book.Publish_date)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2006-January-02"
		date, err = time.Parse(format, book.Publish_date)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "02 Jan 2006"
		date, err = time.Parse(format, book.Publish_date)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2 Jan 2006"
		date, err = time.Parse(format, book.Publish_date)
		if err == nil {return date, nil}
	}

	return date, err
}

func InsertBook(book structs.Book, authorId int) error {
	insertQuery := `
	INSERT INTO books
	(name, author_id, published_date, num_pages, synopsis, ol_id)
	VALUES
	(?, ?, ?, ?, ?, ?)
	`

	bookDate, err := parseBookDateString(book)
	if err != nil {
		return err
	}

	_, err = util.DB.Exec(
		insertQuery,
		book.Title,
		authorId,
		bookDate.Format("2006-01-02"),
		book.Number_of_pages,
		book.Description,
		book.Key,
	)

	if err != nil {
		return err
	}
	return nil
}

func updateBookSynopsis (bookId int, synopsis string) error {
	updateQuery := `
	UPDATE books
	SET synopsis=?
	WHERE id=?;`

	_, err := util.DB.Exec(updateQuery, synopsis, bookId)

	if err != nil {
		return err
	}
	return nil
}
