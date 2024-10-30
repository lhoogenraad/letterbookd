package bookupload

import (
	"fmt"
	"scripts/authors"
)

/**
Will upload books which have authors contained in our
database. These author IDs are grabbed from authors package
*/
func ReadAndUpload() error {
	authorMap, err := authors.GetAllAuthorIds()
	if err != nil {
		return err
	}

	fmt.Println(authorMap)
	return nil
}
