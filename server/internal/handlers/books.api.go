package handlers

import (
	"net/http"
	"encoding/json"
	"server/api"
	"time"
	"server/internal/models"
	
	log "github.com/sirupsen/logrus"
)


func GetBooks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	title := "A Song of Ice and Fire"
	author := "George R.R. Martin"
	pub := time.Date(1995, 1, 1, 0, 0, 0, 0, time.Local)

	book := models.BookData {
		Title: title,
		Author: author,
		Published: pub,
	}

	err := json.NewEncoder(w).Encode(book)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
