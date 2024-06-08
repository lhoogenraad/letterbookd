package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"

	"server/api"
	"server/internal/resources"
	"server/internal/utils"
	// "server/internal/models"
	
	log "github.com/sirupsen/logrus"
)


func CreateReview (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request resources.CreateReviewBody
	claims, isErr := utils.GetClaims(r)
	fmt.Println(claims["userid"])
	fmt.Println(isErr)

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Error(err)
		api.CustomErrorHandler(w, 400, `Invalid review creation body received`)
	}


	// err, code := models.CreateReview(
	// 	userId,
	// 	request.BookId,
	// 	request.Rating,
	// 	request.Description,
	// )

	if err != nil {
		api.CustomErrorHandler(w, 200, fmt.Sprint(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted) // For demo purposes
	json.NewEncoder(w).Encode(`New review created successfully`)
}
