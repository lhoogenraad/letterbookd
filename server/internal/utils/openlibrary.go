package utils

import (
	"fmt"
	"io"
	"net/http"
	"server/internal/resources"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var OPEN_LIBRARY_EDITION_URL = "https://openlibrary.org/search.json?title="

func generateEditionSearchURL (searchValue string) string {
	url := OPEN_LIBRARY_EDITION_URL + searchValue + ""
	url = strings.ReplaceAll(url, " ", "+")
	fmt.Println(url)
	return url
}

func SearchOpenLibrary (search string) (resources.BookData, error) {
	var book resources.BookData
	book, err := queryOpenLibraryForFirstBook(search)
	if err != nil {return book, err}
	return book, nil
}


func queryOpenLibraryForFirstBook (search string) (resources.BookData, error) {
	var firstBook resources.BookData
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
	var parsed resources.OpenLibraryEditionResponse
	err = StringToStruct(sb, &parsed)
	if err != nil {return firstBook, err}
	fmt.Println(parsed.Docs[len(parsed.Docs)-1])
	fmt.Println(sb)
	converted := convertOpenLibaryEditionToBook(parsed.Docs[0])
	fmt.Println(parsed.Docs[0])
	fmt.Println("converted:", converted)
	// return parsed.Docs[0], nil
	return firstBook, nil
}

func convertOpenLibaryEditionToBook(res resources.OpenLibraryEdition) resources.BookData {
	var parsedBook resources.BookData
	parsedBook.Title = res.Title
	parsedBook.Author = res.Author_Name[0]
	pub, err := parseEditionPublishedDateString(res.PublishDate[0])
	if err != nil {
		fmt.Println("We fucked up the parsing", res.PublishDate[0])
	} else { parsedBook.Published = pub }
	// parsedBook.Published = res.PublishDate[0]
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
