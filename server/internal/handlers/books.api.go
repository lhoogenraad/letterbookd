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

	var books []resources.BookData
	var err error
	books, err = models.GetBooks()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	err = json.NewEncoder(w).Encode(books)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
