package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"

	"server/api"
	"server/internal/resources"
	"server/internal/utils"
	"server/internal/models"
	
	log "github.com/sirupsen/logrus"
)



func GetReviewComments (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	//Grab review id param from url
	reviewIdParam := utils.GetParam(r, "reviewId")
	reviewId, err := strconv.Atoi(reviewIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid review ID was given as a parameter.")
		return
	}

	comments, err := models.GetReviewComments(reviewId)


	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}


	// Grab user Id from claims
	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}
	userId := int(claims["userid"].(float64))

	// For each review, check if the owner is the current requester.
	// If they are, set review.OwnedBy to true!
	for i := range comments {
		if comments[i].UserId == userId {
			comments[i].OwnedBy = true
		}
	}

	err = json.NewEncoder(w).Encode(comments)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}


func CreateReviewComment (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}

	//Grab request body
	var request resources.CreateReviewCommentBody
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid request body. Please try again")
		return
	}

	//Convert userId to int
	userId := int(claims["userid"].(float64))

	//Grab review id param from url
	reviewIdParam := utils.GetParam(r, "reviewId")
	reviewId, err := strconv.Atoi(reviewIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid review ID was given as a parameter.")
		return
	}

	err, code := models.CreateReviewComment(reviewId, userId, request)

	if err != nil {
		api.CustomErrorHandler(w, code, fmt.Sprint(err))
		return
	}

	w.WriteHeader(http.StatusAccepted) 
	json.NewEncoder(w).Encode(`Comment created successfully`)
}



func DeleteReviewComment (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}

	//Grab request body
	var request resources.CreateReviewCommentBody
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid request body. Please try again")
		return
	}

	//Convert userId to int
	userId := int(claims["userid"].(float64))

	//Grab review id param from url
	reviewIdParam := utils.GetParam(r, "reviewId")
	reviewId, err := strconv.Atoi(reviewIdParam)

	//Grab review id param from url
	commentIdParam := utils.GetParam(r, "commentId")
	commentId, err := strconv.Atoi(commentIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid review ID was given as a parameter.")
		return
	}

	err, code := models.DeleteReviewComment(reviewId, userId, commentId)

	if err != nil {
		api.CustomErrorHandler(w, code, fmt.Sprint(err))
		return
	}

	w.WriteHeader(http.StatusAccepted) 
	json.NewEncoder(w).Encode(`Archived comment successfully`)
}
