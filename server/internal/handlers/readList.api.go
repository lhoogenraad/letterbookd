package handlers

import (
	"encoding/json"
	"net/http"
	"server/api"
	"server/internal/models"
	"server/internal/resources"
	"server/internal/utils"
	log "github.com/sirupsen/logrus"
)


func GetReadListItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}

	//Convert userId to int
	userId := int(claims["userid"].(float64))

	var readListItems []resources.ReadListItem
	var err error
	readListItems, err = models.GetReadListItems(userId)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	err = json.NewEncoder(w).Encode(readListItems)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}


func AddBookToReadList(w http.ResponseWriter, r *http.Request) {
}
