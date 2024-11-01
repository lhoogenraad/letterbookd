package covers

import (
	"fmt"
	"scripts/books"
)

func AddCoversToBooks() error {
	// bookMap, err := books.GetBookOpenLibIdMap()
	// if err != nil {return err}
	books, err := books.GetAllBooks(1*1000*1)
	if err != nil {return err}
	fmt.Println(books)
	return nil
}
