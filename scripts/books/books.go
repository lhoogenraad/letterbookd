package books

import (
	"encoding/json"
	"scripts/util"
	"strings"
	"fmt"
)

type Book struct{
	Languages []struct  {
		Key string `json: "key"`
	}
	Title string `json:"title"`
	Number_of_pages int16 `json:"number_of_pages`
	Publish_date string `json:"publish_date"`
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

func hasEnoughPages(book Book) bool {
	num_pages := book.Number_of_pages
	return num_pages > 200 
}

func ReadAndUpload () error {
	filepath := `/home/leon/Downloads/ol_dump_editions_2024-09-30.txt`
	scanner, err := util.GetScanner(filepath)

	if err != nil{
		fmt.Println(err)
		return err
	}

	i := 0
	max := 100000
	for scanner.Scan() && i < max {
		book := getLineAsJSON(scanner.Text())
		if hasEnoughPages(book) {
			fmt.Println(book.Title, book.Publish_date)
		}
		i++
	}

	return nil
}

