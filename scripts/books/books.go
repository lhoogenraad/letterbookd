package books

import (
	"encoding/json"
	"scripts/util"
	"strings"
	"fmt"
	"time"
)

type Book struct{
	Languages []struct  {
		Key string `json: "key"`
	}
	Title string `json:"title"`
	Number_of_pages int16 `json:"number_of_pages`
	Publish_date string `json:"publish_date"`
	Authors []struct {
		Key string `json: "key"`
	}
}

func getLineAsJSON (text string) Book {
	book := Book{}
	textSplit := strings.Split(text, "{")
	textSplit = textSplit[1:]
cleaned := strings.Join(textSplit, "{")
	cleaned = "{" + cleaned
	json.Unmarshal([]byte(cleaned), &book)

	return book
}

func isEnglish(book Book) bool {
	languages := book.Languages
	for _, val := range languages {
		if val.Key == "/languages/eng"{
			return true
		}
	}
	return false
}

func isRecent(book Book) bool {
	var format string
	var dateString string = book.Publish_date
	var date time.Time

	// Commonly the Pub dates are just the year
	if len(dateString) == 4 {
		format = "2000"
	} else {
		format = "Jan 02, 2006"
	}

	date, err := time.Parse(format, dateString)

	minDate, err := time.Parse("2006-02-01", "1980-01-01")
	if err == nil && date.After(minDate){
		return true
	} 
	return false
}

func hasEnoughPages(book Book) bool {
	num_pages := book.Number_of_pages
	return num_pages > 200 
}

func shouldAddBook(book Book) bool {
	return hasEnoughPages(book) && isRecent(book) && isEnglish(book)
}

func GetValidBooks () ([]Book, error) {
	filepath := `/home/leon/Downloads/ol_dump_editions_2024-09-30.txt`
	fmt.Println(`Retrieving valid books from`, filepath)
	scanner, err := util.GetScanner(filepath)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}


	i := 0
	max := 1000000
	var validBooks []Book
	for scanner.Scan() && i < max{
		book := getLineAsJSON(scanner.Text())
		if shouldAddBook(book) {
			validBooks = append(validBooks, book)
		}
		if i % 10000 == 0{
			fmt.Println(`Got to iteration`, i)
		}
		i++
	}
	err = scanner.Err()
	if err != nil{
		fmt.Println(`\n\nEncountered err:`, err, "\n\n")
	}
	
	fmt.Println(`Found`, len(validBooks), `valid books`)
	return validBooks, nil
}

