package covers

import (
	"scripts/books"
	"strconv"
	"fmt"
)


//creating a function to add zeroes to a string
func PadLeft(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}

func AddCoversToBooks() error {
	bookMap, err := books.GetBookOpenLibIdMap()
	if err != nil {return err}
	books, err := books.GetAllBooks(1000*1000*1)
	if err != nil {return err}
	for _, book := range books {
		for _, cover := range book.Covers {
			filepath := getFilepathOfCoverInt(cover)
			bookId, ok := bookMap[book.Key]
			if !ok {
				fmt.Println(book.Title, " was not found in map")
				continue
			}
			url, err := UploadCoverAndGetURL(filepath)
			if err == nil {
				fmt.Printf("Setting book id %d (%s) (%s) cover url to %s \n", bookId, book.Key, book.Title, url)
				SetBookUrl(bookId, url)
			} else {
				fmt.Println("Failed to upload book cover", filepath, "for book", book.Title)
			}
		}
	}
	return nil
}

func getFilepathOfCoverInt(cover int) string {
	filepath := `/home/leon/Documents/letterbookd_files/covers_unzipped/` + determineFilename(cover)
	return filepath
}

func determineFilename(cover int) string {
	coverStr := strconv.Itoa(cover)
	filename := PadLeft(coverStr, 10) + ".jpg"
	return filename
}
