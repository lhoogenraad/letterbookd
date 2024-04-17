package handlers

import (
	"net/http"
	"encoding/json"
	"server/api"
	"fmt"
	"server/internal/resources"
	"server/internal/utils"
	
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

	fmt.Println(request)
	var hash string 
	hash, err = utils.HashPassword(request.Password)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	fmt.Println("hashed pword: " + hash)
}
