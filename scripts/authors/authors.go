package authors

import (
	"encoding/json"
	"scripts/util"
	"fmt"
	"scripts/books"
	"time"
	"strings"
)


type Author struct {
	Key string `json: "key"`
	Name string `json: "name"`
	Birth_date string `json: "birth_date"`
	DOB time.Time
}


func getLineAsJSON (text string) Author {
	textSplit := strings.Split(text, "{")
	textSplit = textSplit[1:]
	cleaned := strings.Join(textSplit, "{")
	cleaned = "{" + cleaned
	var author Author
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

func hasBirthDate(author Author) ( time.Time, bool ) {
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

	var authorsToAdd []Author
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
