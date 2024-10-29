package authors

import (
	"encoding/json"
	// "scripts/util"
	"fmt"
	"scripts/books"
	"time"
	"scripts/util"
	"strings"
)


type Author struct {
	Key string `json: "key"`
	Name string `json: "name"`
	Birth_date string `json: "birth_date"`
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
	date_of_birth, err := time.Parse("2006", dob)
	if err == nil {return date_of_birth, true}
	date_of_birth, err = time.Parse("01 Feb, 2006", dob)
	if err == nil {return date_of_birth, true}
	date_of_birth, err = time.Parse("2006-02-01", dob)
	if err != nil {return date_of_birth, false}
	return date_of_birth, true;
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
				fmt.Println(author, "dob:", dob)
				authorsToAdd = append(authorsToAdd, author)
				delete(authorIdMap, id)
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
