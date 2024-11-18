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


func CreateReview (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request resources.CreateReviewBody
	claims, ok := utils.GetClaims(r)

	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Error(err)
		api.CustomErrorHandler(w, 400, `Invalid review creation body received`)
	}

	//Convert userId to int
	userId := int(claims["userid"].(float64))
	bookIdParam := utils.GetParam(r, "bookId")
	bookId, err := strconv.Atoi(bookIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid book ID was given as a parameter.")
		return
	}
	err, code := models.CreateReview(userId, bookId, request)

	if err != nil {
		api.CustomErrorHandler(w, code, fmt.Sprint(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted) // For demo purposes
	json.NewEncoder(w).Encode(`New review created successfully`)
}


func UpdateReview (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request resources.UpdateReviewBody
	claims, ok := utils.GetClaims(r)

	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Error(err)
		api.CustomErrorHandler(w, 400, `Invalid review creation body received`)
	}

	//Convert userId to int
	userId := int(claims["userid"].(float64))
	reviewIdParam := utils.GetParam(r, "reviewId")
	reviewId, err := strconv.Atoi(reviewIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid review ID was given as a parameter.")
		return
	}

	err, code := models.UpdateReview(userId, reviewId, request)

	if err != nil {
		api.CustomErrorHandler(w, code, fmt.Sprint(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted) 
	json.NewEncoder(w).Encode(`Review updated successfully`)
}


func GetBookReviews (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Grab user Id from claims
	claims, ok := utils.GetClaims(r)
	if !ok {
		log.Error("Something went wrong grabbing token claim info")
		api.InternalErrorHandler(w)
	}
	userId := int(claims["userid"].(float64))

	bookIdParam := utils.GetParam(r, "bookId")
	bookId, err := strconv.Atoi(bookIdParam)

	if err != nil {
		api.CustomErrorHandler(w, 400, "Invalid book ID was given as a parameter.")
		return
	}

	reviews, err := models.GetBookReviews(bookId, userId)

	// For each review, check if the owner is the current requester.
	// If they are, set review.OwnedBy to true!
	for i := range reviews {
		fmt.Print(reviews[i])
		if reviews[i].UserId == userId {
			reviews[i].OwnedBy = true
		}
	}

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	err = json.NewEncoder(w).Encode(reviews)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
