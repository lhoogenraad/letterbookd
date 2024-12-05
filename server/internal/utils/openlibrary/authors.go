package openlibrary

import (
	"fmt"
	"io"
	"net/http"
	"github.com/lhoogenraad/letterbookd/internal/resources"
	"github.com/lhoogenraad/letterbookd/internal/utils"
	"github.com/lhoogenraad/letterbookd/internal/tools"
	"strings"
)

var OPEN_LIBRARY_AUTHOR_SEARCH_URL = "https://openlibrary.org/authors/"

func GetAuthorId (authorOlId string) (int, error) {
	var author resources.Author
	fmt.Println(authorOlId)
	fmt.Println(1)
	// Check if author already exists based on given author Open lib ID key
	author, exists, _ := getAuthorFromDB(authorOlId)
	if exists == true {	return author.Id, nil } 

	fmt.Println(2)
	author, err := getAuthorFromOL(authorOlId)
	if err != nil {return author.Id, err}

	fmt.Println(3)
	// Insert author and retrive it's ID
	err = saveAuthor(author, authorOlId)
	if err != nil {return author.Id, err}
	fmt.Println(4)
	author, _, err = getAuthorFromDB(authorOlId)
	if err != nil {return author.Id, err}
	fmt.Println(5)
	return author.Id, nil
}

func getAuthorFromDB(authorOlId string) (resources.Author, bool, error) {
	fmt.Println("Author ID:", authorOlId)
	var author resources.Author
	query := `SELECT id FROM authors WHERE ol_id LIKE ?`
	filter := "%" + authorOlId + "%"

	row := tools.DB.QueryRow(query, filter)
	err := row.Scan(&author.Id)
	if err != nil {return author, false, err}

	return author, true, nil
}

func saveAuthor(author resources.Author, authorOlId string) error {
	insertQuery := `
	INSERT INTO authors
	(first_name, last_name, date_of_birth, ol_id)
	VALUES
	(?, ?, ?, ?)
	`
	_, err := tools.DB.Exec(
		insertQuery, 
		author.FirstName, 
		author.LastName, 
		author.DateOfBirth.Format("2006-01-02"),
		"/authors/" + authorOlId,
	)

	return err
}

func getAuthorFromOL (authorOlId string) (resources.Author, error) {
	var author resources.Author
	resp, err := http.Get(getAuthorSearchURL(authorOlId))
	if err != nil {return author, err}

	body, err := io.ReadAll(resp.Body)
	if err != nil {return author, err}

	sb := string(body)
	var parsedAuthor resources.AuthorOL
	err = utils.StringToStruct(sb, &parsedAuthor)
	if err != nil {return author, err}

	author = convertAuthorOLToAuthor(parsedAuthor)
	return author, err
}

func getAuthorSearchURL (authorOlId string) string {
	return OPEN_LIBRARY_AUTHOR_SEARCH_URL + authorOlId + ".json"
}

func convertAuthorOLToAuthor(authorOl resources.AuthorOL) resources.Author {
	var author resources.Author

	author.FirstName, author.LastName = splitAuthorFullname(authorOl.Name)

	var err error
	author.DateOfBirth, err = utils.ParseStringToTime(authorOl.Birth_Date)
	if err != nil {fmt.Println("Error, Couldn't parse author DOB.", authorOl.Birth_Date, err)}

	return author
}

/**
Splits given author name at spaces.
First name is considered split[0]
Last name is considered the rest of the split
*/
func splitAuthorFullname (name string) (string, string) {
	var firstName, lastName string

	split := strings.Split(name, " ")

	firstName = split[0]
	lastName = strings.Join(split[1:], " ")

	return firstName, lastName
}
