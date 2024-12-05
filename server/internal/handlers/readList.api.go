package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/lhoogenraad/letterbookd/api"
	"github.com/lhoogenraad/letterbookd/internal/models"
	"github.com/lhoogenraad/letterbookd/internal/resources"
	"github.com/lhoogenraad/letterbookd/internal/utils"
	"fmt"
	"strconv"
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
	w.Header().Set("Content-Type", "application/json")

	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}

	//Convert userId to int
	userId := int(claims["userid"].(float64))

	bookIdParam := utils.GetParam(r, "bookId")
	bookId, err := strconv.Atoi(bookIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid book ID was given as a parameter.")
		return
	}

	//Grab request body
	var request resources.ReadListModReq
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid request body. Please try again")
		return
	}

	err, code := models.AddBookToReadlist(bookId, userId, request)

	if err != nil {
		api.CustomErrorHandler(w, code, fmt.Sprint(err))
		return
	}

	w.WriteHeader(http.StatusAccepted) 
	json.NewEncoder(w).Encode(`We've added this book to your readlist!`)
}

func UpdateReadListItem (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}

	//Convert userId to int
	userId := int(claims["userid"].(float64))

	bookIdParam := utils.GetParam(r, "bookId")
	bookId, err := strconv.Atoi(bookIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid book ID was given as a parameter.")
		return
	}

	//Grab request body
	var request resources.ReadListModReq
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid request body. Please try again")
		return
	}

	err, code := models.UpdateReadListItem(bookId, userId, request)

	if err != nil {
		api.CustomErrorHandler(w, code, fmt.Sprint(err))
		return
	}

	w.WriteHeader(http.StatusAccepted) 
	json.NewEncoder(w).Encode(`We've updated this book successfully!`)
}


func DeleteReadListItem (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}

	//Convert userId to int
	userId := int(claims["userid"].(float64))

	bookIdParam := utils.GetParam(r, "bookId")
	bookId, err := strconv.Atoi(bookIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid book ID was given as a parameter.")
		return
	}

	err, code := models.DeleteReadListItem(bookId, userId)

	if err != nil {
		api.CustomErrorHandler(w, code, fmt.Sprint(err))
		return
	}

	w.WriteHeader(http.StatusAccepted) 
	json.NewEncoder(w).Encode(`We've removed this item from your readlist successfully!`)
}
