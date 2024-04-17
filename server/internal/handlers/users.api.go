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


func Signup (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
