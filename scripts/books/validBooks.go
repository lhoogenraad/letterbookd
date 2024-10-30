package books

import (
	"encoding/json"
	"scripts/util"
	"strings"
	"fmt"
	"time"
	"scripts/structs"
)


func getLineAsJSON (text string) structs.Book {
	book := structs.Book{}
	textSplit := strings.Split(text, "{")
	textSplit = textSplit[1:]
	cleaned := strings.Join(textSplit, "{")
	cleaned = "{" + cleaned
	fmt.Println(cleaned)
	json.Unmarshal([]byte(cleaned), &book)

	return book
}

func isEnglish(book structs.Book) bool {
	languages := book.Languages
	for _, val := range languages {
		if val.Key == "/languages/eng"{
			return true
		}
	}
	return false
}

func isRecent(book structs.Book) bool {
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

func hasEnoughPages(book structs.Book) bool {
	num_pages := book.Number_of_pages
	return num_pages > 200 
}

func shouldAddBook(book structs.Book) bool {
	return hasEnoughPages(book) && isRecent(book) && isEnglish(book)
}

func GetValidBooks () ([]structs.Book, error) {
	filepath := `/home/leon/Downloads/ol_dump_editions_2024-09-30.txt`
	fmt.Println(`Retrieving valid books from`, filepath)
	scanner, err := util.GetScanner(filepath)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}


	i := 0
	max := 100000000
	var validBooks []structs.Book
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


func GetAllBooks (limit int) ([]structs.Book, error) {
	filepath := `/home/leon/Downloads/ol_dump_editions_2024-09-30.txt`
	fmt.Println(`Retrieving all books from`, filepath, `up until`, limit)
	scanner, err := util.GetScanner(filepath)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	i := 0
	var books []structs.Book
	for scanner.Scan() && i < limit{
		book := getLineAsJSON(scanner.Text())
		books = append(books, book)
		if i % 10000 == 0{
			fmt.Println(`Got to iteration`, i)
		}
		i++
	}
	err = scanner.Err()
	if err != nil{
		fmt.Println(`\n\nEncountered err:`, err, "\n\n")
	}

	return books, nil
}


