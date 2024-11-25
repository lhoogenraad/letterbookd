package openlibrary

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/internal/models"
	"server/internal/resources"
	"server/internal/tools"
	"server/internal/utils"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var OPEN_LIBRARY_EDITION_URL = "https://openlibrary.org/search.json?"
var OPEN_LIBRARY_COVER_URL = "https://covers.openlibrary.org/b/olid/"
var OPEN_LIBRARY_OLID_SEARCH_URL = "https://openlibrary.org/works/"

func generateEditionSearchURL (title string, author string, publisher string) string {
	req, err := http.NewRequest("GET", OPEN_LIBRARY_EDITION_URL, nil)
	if err != nil {log.Print(err)}
	
	filter := req.URL.Query()
	if title != "" {
		filter.Add("title", title)
	}

	if author != "" {
		filter.Add("author", author)
	}

	if publisher != "" {
		filter.Add("publisher", publisher)
	}
    req.URL.RawQuery = filter.Encode()

	return req.URL.String()
}

func generateCoverSearchURL (coverString string) string {
	searchURL := OPEN_LIBRARY_COVER_URL + coverString + "-L.jpg"
	return searchURL
}

func generateOLIDSearchURL (olId string) string {
	searchURL := OPEN_LIBRARY_OLID_SEARCH_URL + olId + ".json"
	return searchURL
}

func retrieveCoverImage(olCoverId string, save bool) (string, error){
	if olCoverId == "" {return "", nil}
	path := "/home/leon/Documents/letterbookd/client/public/covers/" + olCoverId + ".jpg"
	url := generateCoverSearchURL(olCoverId)
	resp, err := http.Get(url)
	if err != nil {return path, err}
	body := resp.Body

	if save {
		fmt.Printf("Uploading a cover url!\nolCoverId: %s\tSave Path: %s\n", olCoverId, path)
		err = saveCoverImage(body, path)
		if err != nil {return path, err}
	}

	if save {
		return "covers/" + olCoverId + ".jpg", nil
	} else {
		return url, nil
	}
}

func SearchOpenLibrary (title string, author string, publisher string) (resources.BookDataOL, error) {
	var book resources.BookDataOL
	book, err := queryOpenLibraryForEditions(title, author, publisher)
	if err != nil {return book, err}

	path, err := retrieveCoverImage(book.CoverEdition, false)
	if err != nil {return book, err}

	book.CoverURL = path
	return book, nil
}


func queryOpenLibraryForEditions (title string, author string, publisher string) (resources.BookDataOL, error) {
	var firstBook resources.BookDataOL
	resp, err := http.Get(generateEditionSearchURL(title, author, publisher))
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
	if err != nil {return firstBook, err}

	firstBook.CoverURL, err = retrieveCoverImage(firstBook.CoverEdition, true)
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
	fmt.Println(sb)
	var parsed resources.OpenLibraryEdition
	err := utils.StringToStruct(sb, &parsed)
	if err != nil {return book, err}

	book = convertOpenLibaryEditionToBook(parsed)
	return book, nil
}

func convertOpenLibaryEditionToBook(res resources.OpenLibraryEdition) resources.BookDataOL {
	var parsedBook resources.BookDataOL
	parsedBook.Title = res.Title
	fmt.Println("1")
	fmt.Println(res)
	fmt.Println(res.AuthorKey, res.Authors)
	// Parse author ID
	if len(res.AuthorKey) > 0 {
		fmt.Println("none")
		parsedBook.Author = res.Author_Name[0]
		parsedBook.AuthorOLId = res.AuthorKey[0]
	} else if len(res.Authors) > 0 {
		fmt.Println("a")
		authorIdPathSplit := strings.Split(res.Authors[0].Author.Key, "/")
		fmt.Println("b")
		parsedBook.AuthorOLId = authorIdPathSplit[len(authorIdPathSplit)-1]
		fmt.Println("c")
	} 

	fmt.Println("2")
	// Parse publish date
	if len(res.PublishDate) > 0{
		pub, err := utils.ParseStringToTime(res.PublishDate[0])
		if err != nil {
			fmt.Println("We messed up the parsing", res.PublishDate[0])
		} else { parsedBook.Published = pub }
	}

	fmt.Println("3")
	// Grab first edition key available
	if len(res.EditionKey) > 0 {
		parsedBook.OpenLibraryKey = res.EditionKey[0]
	}

	fmt.Println("4")
	splitWorkID := strings.Split(res.OlID, "/")
	parsedBook.OlID = splitWorkID[len(splitWorkID)-1]
	fmt.Println("res.Covers:", res.Covers)
	if len(res.Covers) > 0 {
		parsedBook.CoverEdition = "OL" + strconv.Itoa(res.Covers[0]) + "M"
	}else if res.Cover_Edition.Key != "" {
		splitCoverEditionKey := strings.Split(res.Cover_Edition.Key, "/")
		parsedBook.CoverEdition = splitCoverEditionKey[ len(splitCoverEditionKey)-1 ]
	} else {
		parsedBook.CoverEdition = res.CoverEditionKey
	}
	parsedBook.Synopsis = res.Description.Value
	fmt.Println("Parsed Book:", parsedBook)
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
	if olIdExists(olId) {return book, errors.New("It appears we already have this book in Shelfd") }
	book, err := searchOpenLibraryForOLID(olId)
	if err != nil {return book, err}

	fmt.Println("Hello?")
	book.AuthorId, err = GetAuthorId(book.AuthorOLId)
	if err != nil {return book, err}

	fmt.Println("Parsed book:", book)
	err = models.UploadBook(book)
	if err != nil {return book, err}

	return book, nil
}


/**
Checks if we currently have a book in our database with a matching
open library ID
*/
func olIdExists (olId string) bool {
	query := `SELECT id FROM books WHERE ol_id LIKE ?`
	filter := "%" + olId + "%"

	var id int
	err := tools.DB.QueryRow(query, filter).Scan(&id)

	if err == sql.ErrNoRows {return false}

	return true
}
