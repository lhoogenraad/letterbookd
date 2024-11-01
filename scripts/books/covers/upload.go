package covers

import (
	"fmt"
	"scripts/books"
	"scripts/util"
)

func AddCoversToBooks() error {
	// bookMap, err := books.GetBookOpenLibIdMap()
	// if err != nil {return err}
	books, err := books.GetAllBooks(1*1000*1)
	if err != nil {return err}
	fmt.Println(books)
	return nil
}

func SetBookUrl(bookId int, coverURL string) error {
	updateQuery := `UPDATE books SET cover_url = ? WHERE id = ?`
	_, err := util.DB.Exec(updateQuery, coverURL, bookId)
	return err
}
