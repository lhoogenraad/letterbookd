package handlers

import (
	"encoding/json"
	"net/http"
	"server/api"
	log "github.com/sirupsen/logrus"
)


func HelloJamie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	msg := "BING BONG"

	err := json.NewEncoder(w).Encode(msg)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
