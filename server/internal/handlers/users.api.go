package handlers

import (
	"net/http"
	"encoding/json"
	"server/api"

	"fmt"
	"server/internal/resources"
	"server/internal/utils"
	"server/internal/models"

	log "github.com/sirupsen/logrus"
)


func Signup (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request resources.SignupRequestBody

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Error(err)
		api.CustomErrorHandler(w, 400, `Invalid signup body received`)
	}

	var hash string 
	hash, err = utils.HashPassword(request.Password)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	err, code := models.AddUser(
		request.Email,
		hash,
		request.FirstName,
		request.LastName,
	)

	if err != nil {
		api.CustomErrorHandler(w, code, fmt.Sprint(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted) // For demo purposes
	json.NewEncoder(w).Encode(`Created new user account for ` + request.Email)
}
