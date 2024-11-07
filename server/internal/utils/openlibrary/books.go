package openlibrary

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/internal/resources"
	"server/internal/utils"
	"strings"

	log "github.com/sirupsen/logrus"
)

var OPEN_LIBRARY_EDITION_URL = "https://openlibrary.org/search.json?q="
var OPEN_LIBRARY_COVER_URL = "https://covers.openlibrary.org/b/olid/"
var OPEN_LIBRARY_OLID_SEARCH_URL = "https://openlibrary.org/works/"

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

func generateOLIDSearchURL (olId string) string {
	url := OPEN_LIBRARY_OLID_SEARCH_URL + olId + ".json"
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
	// err = utils.SaveBook(book)
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
		log.Error(err)
		return firstBook, err
	}

	firstBook, err = parseOLServerResponse(body)
	if err != nil {
		log.Error(err)
		return firstBook, err
	}
	// return parsed.Docs[0], nil
	return firstBook, nil
}


func searchOpenLibraryForOLID (olId string) (resources.BookDataOL, error) {
	var firstBook resources.BookDataOL
	resp, err := http.Get(generateOLIDSearchURL(olId))
	if err != nil {
		return firstBook, err
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return firstBook, err
	}

	firstBook, err = parseOLWorksServerResponse(body)
	if err != nil {
		log.Error(err)
		return firstBook, err
	}
	// return parsed.Docs[0], nil
	return firstBook, nil
}


func parseOLServerResponse (body []byte) (resources.BookDataOL, error) {
	var firstBook resources.BookDataOL

	sb := string(body)
	var parsed resources.OpenLibraryEditionResponse
	err := utils.StringToStruct(sb, &parsed)
	if err != nil {return firstBook, err}

	if len(parsed.Docs) < 1 {
		return firstBook, errors.New("No books found sorry")
	}
	firstBook = convertOpenLibaryEditionToBook(parsed.Docs[0])
	return firstBook, nil
}

func parseOLWorksServerResponse (body []byte) (resources.BookDataOL, error) {
	var book resources.BookDataOL

	sb := string(body)
	var parsed resources.OpenLibraryEdition
	err := utils.StringToStruct(sb, &parsed)
	if err != nil {return book, err}

	book = convertOpenLibaryEditionToBook(parsed)
	return book, nil
}

func convertOpenLibaryEditionToBook(res resources.OpenLibraryEdition) resources.BookDataOL {
	var parsedBook resources.BookDataOL
	parsedBook.Title = res.Title
	// Parse author ID
	if len(res.AuthorKey) > 0 {
		parsedBook.Author = res.Author_Name[0]
		parsedBook.AuthorOLId = res.AuthorKey[0]
	} else if len(res.Authors) > 0 {
		authorIdPathSplit := strings.Split(res.Authors[0].Author.Key, "/")
		parsedBook.AuthorOLId = authorIdPathSplit[len(authorIdPathSplit)-1]
	}

	// Parse publish date
	if len(res.PublishDate) > 0{
		pub, err := utils.ParseStringToTime(res.PublishDate[0])
		if err != nil {
			fmt.Println("We messed up the parsing", res.PublishDate[0])
		} else { parsedBook.Published = pub }
	}

	// Grab first edition key available
	if len(res.EditionKey) > 0 {
		parsedBook.OpenLibraryKey = res.EditionKey[0]
	}
	splitWorkID := strings.Split(res.OlID, "/")
	parsedBook.OlID = splitWorkID[len(splitWorkID)-1]
	parsedBook.CoverEdition = res.CoverEditionKey
	parsedBook.Synopsis = res.Description
	return parsedBook
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
	book, err := searchOpenLibraryForOLID(olId)
	if err != nil {return book, err}

	authorId, err := GetAuthorId(book.AuthorOLId)
	if err != nil {return book, err}
	fmt.Println("author Id:", authorId)

	// fmt.Println("Book:", book)
	return book, nil
}


