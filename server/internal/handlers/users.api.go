package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/lhoogenraad/letterbookd/api"

	"fmt"
	"github.com/lhoogenraad/letterbookd/internal/resources"
	"github.com/lhoogenraad/letterbookd/internal/utils"
	"github.com/lhoogenraad/letterbookd/internal/models"

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



func Signin (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request resources.SigninRequestBody

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Error(err)
		api.CustomErrorHandler(w, 400, `Invalid sign in body received`)
	}

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	var user resources.User
	user, err, code := models.GetUser(request.Email)

	if err != nil {
		api.CustomErrorHandler(w, code, fmt.Sprint(err))
		return
	}

	// Password validation
	var validPassword bool = utils.CheckPasswordHash(request.Password, user.PasswordHash)
	if !validPassword {
		api.CustomErrorHandler(w, 401, `Invalid email or password. Please try again.`)
		return
	}

	// Token generation
	token, err := utils.GenerateToken(
		uint(user.Id),
		user.FirstName,
		user.Email,
	)

	if err != nil {
		api.CustomErrorHandler(w, 500, `Sorry, we couldn't generate an auth token for some reason. Please contact us if this issue persits`)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted) // For demo purposes
	json.NewEncoder(w).Encode(token)
}
