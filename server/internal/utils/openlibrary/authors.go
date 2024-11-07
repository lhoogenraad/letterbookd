package openlibrary

import (
	"fmt"
	"io"
	"net/http"
	"server/internal/resources"
	"server/internal/utils"
	"strings"
)

var OPEN_LIBRARY_AUTHOR_SEARCH_URL = "https://openlibrary.org/authors/"

func RetrieveAuthorFromOL (authorId string) (resources.Author, error) {
	var author resources.Author

	resp, err := http.Get(getAuthorSearchURL(authorId))
	if err != nil {return author, err}

	body, err := io.ReadAll(resp.Body)
	if err != nil {return author, err}

	sb := string(body)
	var parsedAuthor resources.AuthorOL
	err = utils.StringToStruct(sb, &parsedAuthor)
	if err != nil {return author, err}

	author = convertAuthorOLToAuthor(parsedAuthor)

	return author, nil
}

func getAuthorSearchURL (authorId string) string {
	return OPEN_LIBRARY_AUTHOR_SEARCH_URL + authorId + ".json"
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
