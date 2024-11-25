package authors

import (
	"encoding/json"
	"scripts/util"
	"fmt"
	"scripts/books"
	"time"
	"strings"
	"scripts/structs"
)




func getLineAsJSON (text string) structs.Author {
	textSplit := strings.Split(text, "{")
	textSplit = textSplit[1:]
	cleaned := strings.Join(textSplit, "{")
	cleaned = "{" + cleaned
	var author structs.Author
	json.Unmarshal([]byte(cleaned), &author)

	return author
}

// Creates map of author IDs to 
func getListOfAuthorIds() (map[string]int, error) {
	var authors  = make(map[string]int)
	books, err:= books.GetValidBooks()
	if err != nil {
		return authors, err
	}

	for _, book := range books {
		bookAuthors := book.Authors
		for _, author := range bookAuthors {
			authorId := author.Key
			if authors[authorId] == 0 {
				fmt.Println(`Found new author ID:`, authorId)
				authors[authorId] = 1
			}
		}
	}
	return authors, nil
}

func hasBirthDate(author structs.Author) ( time.Time, bool ) {
	dob := author.Birth_date
	var format string
	if len(dob) == 4 {
		format = "2006"
	} else if len(dob) == 10 {
		format = "2006-01-02"
	} else {
		format = "2 January 2006"
	}
	date_of_birth, err := time.Parse(format, dob)
	if err != nil {
		fmt.Println(err)
	}

	return date_of_birth, err == nil;
}

func ReadAndUpload () error {
	authorIdMap, err := getListOfAuthorIds()
	if err != nil {
		fmt.Println(err)
		return err
	}

	filepath := `/home/leon/Downloads/ol_dump_authors_2024-09-30.txt`
	scanner, err := util.GetScanner(filepath)

	if err != nil{
		fmt.Println(err)
		return err
	}

	var authorsToAdd []structs.Author
	i := 0
	for scanner.Scan() {
		author := getLineAsJSON(scanner.Text())
		id := author.Key
		_, exists := authorIdMap[id]

		if exists && author.Birth_date != ""{
			dob, ok := hasBirthDate(author)
			if ok {
				author.DOB = dob
				fmt.Println(`Adding author`, author)
				err := UploadAuthor(author)
				if err != nil {
					delete(authorIdMap, id)
				} else {
					fmt.Println(`Failed to upload author `, author, "\n\n")
				}
			}
		}
		if i % 1000000 == 0{
			fmt.Println(authorIdMap[id])
			fmt.Println(`Got to iteration`, i)
		}
		i++
	}
	fmt.Println(`authors to add:`, authorsToAdd)
	return nil
}

/**
Retrieves all author IDs and puts them into a map[string]int
The string key is the ol ID, and the int value is Shelfd's author ID
*/
func GetAllAuthorIds () (map[string]int, error) {
	var authors = make(map[string]int)
	var selectQuery string = `SELECT id, ol_id FROM authors WHERE ol_id IS NOT NULL;`
	rows, err := util.DB.Query(selectQuery);

	if err != nil {return authors, err}

	for rows.Next() {
		var olId string
		var authorId int
		err := rows.Scan(&authorId, &olId)
		if err != nil {return authors, err}
		
		authors[olId] = authorId
	}

	return authors, nil
}
