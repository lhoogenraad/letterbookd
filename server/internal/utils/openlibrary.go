package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/internal/resources"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var OPEN_LIBRARY_EDITION_URL = "https://openlibrary.org/search.json?q="
var OPEN_LIBRARY_COVER_URL = "https://covers.openlibrary.org/b/olid/"

func generateEditionSearchURL (searchValue string) string {
	url := OPEN_LIBRARY_EDITION_URL + searchValue + "&limit=1"
	url = strings.ReplaceAll(url, " ", "+")
	return url
}

func generateCoverSearchURL (coverString string) string {
	url := OPEN_LIBRARY_COVER_URL + coverString + "-L.jpg"
	url = strings.ReplaceAll(url, " ", "+")
	return url
}

func retrieveAndSaveCoverImage(olCoverId string) (string, error){
	if olCoverId == "" {return "", nil}
	path := "/home/leon/Documents/letterbookd/client/public/covers/" + olCoverId + ".jpg"
	fmt.Printf("Uploading a cover url!\nolCoverId: %s\tSave Path: %s\n", olCoverId, path)
	url := generateCoverSearchURL(olCoverId)
	resp, err := http.Get(url)
	if err != nil {return path, err}
	body := resp.Body

	err = saveCoverImage(body, path)
	if err != nil {return path, err}

	return "covers/" + olCoverId + ".jpg", nil
}

func SearchOpenLibrary (search string) (resources.BookDataOL, error) {
	var book resources.BookDataOL
	book, err := queryOpenLibraryForFirstBook(search)
	if err != nil {return book, err}

	path, err := retrieveAndSaveCoverImage(book.CoverEdition)
	if err != nil {return book, err}

	book.CoverURL = path
	err = SaveBook(book)
	if err != nil {return book, err}

	return book, nil
}


func queryOpenLibraryForFirstBook (search string) (resources.BookDataOL, error) {
	var firstBook resources.BookDataOL
	resp, err := http.Get(generateEditionSearchURL(search))
	if err != nil {
		return firstBook, err
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Println(sb)
	var parsed resources.OpenLibraryEditionResponse
	err = StringToStruct(sb, &parsed)
	if err != nil {return firstBook, err}

	if len(parsed.Docs) < 1 {
		return firstBook, errors.New("No books found for search " + search)
	}

	firstBook = convertOpenLibaryEditionToBook(parsed.Docs[0])
	// return parsed.Docs[0], nil
	return firstBook, nil
}

func convertOpenLibaryEditionToBook(res resources.OpenLibraryEdition) resources.BookDataOL {
	var parsedBook resources.BookDataOL
	parsedBook.Title = res.Title
	if len(res.AuthorKey) > 0 {
		parsedBook.Author = res.Author_Name[0]
		parsedBook.AuthorOLId = res.AuthorKey[0]
	}
	if len(res.PublishDate) > 0{
		pub, err := parseEditionPublishedDateString(res.PublishDate[0])
		if err != nil {
			fmt.Println("We messed up the parsing", res.PublishDate[0])
		} else { parsedBook.Published = pub }
	}
	if len(res.EditionKey) > 0 {
		parsedBook.OpenLibraryKey = res.EditionKey[0]
	}
	parsedBook.CoverEdition = res.CoverEditionKey
	return parsedBook
}

func parseEditionPublishedDateString(dateString string)  ( time.Time, error ) {
	var format string
	var date time.Time

	// Commonly the Pub dates are just the year
	if len(dateString) == 4 {
		format = "2006"
	} else {
		format = "January 02, 2006"
	}

	date, err := time.Parse(format, dateString)

	if err != nil {
		format = "January 2, 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "January 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "Jan 2, 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "Jan 02, 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2006-01-02"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2006-Jan-02"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2006-January-02"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "02 Jan 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2 Jan 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	return date, err
}


func saveCoverImage(stream io.Reader, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {return err}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {return err}

	return nil
}



func UploadBookFromOpenLibrary (olId string) (resources.BookDataOL, error) {
	var book resources.BookDataOL
	book, err := queryOpenLibraryForFirstBook(olId)
	fmt.Println(err, book)
	if err != nil {return book, nil}

	fmt.Println("Book:", book)
	return book, nil
}
