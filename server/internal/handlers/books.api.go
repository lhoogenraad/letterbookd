package handlers

import (
	"net/http"
	"encoding/json"
	"server/api"
	// "time"
	"server/internal/resources"
	"server/internal/models"
	
	log "github.com/sirupsen/logrus"
)


func GetBooks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// title := "A Song of Ice and Fire"
	// author := "George R.R. Martin"
	// pub := time.Date(1995, 1, 1, 0, 0, 0, 0, time.Local)

	var books []resources.BookData
	var err error
	books, err = models.GetBooks()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// book := resources.BookData {
	// 	Title: title,
	// 	Author: author,
	// 	Published: pub,
	// 	NumPages: 153,
	// 	CoverURL: "http://www.google.com",
	// }

	err = json.NewEncoder(w).Encode(books)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
