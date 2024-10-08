package api

import (
	"encoding/json"
	"net/http"
	log "github.com/sirupsen/logrus"
)

type Error struct {
	Code int
	Message string
}


func writeError(w http.ResponseWriter, message string, code int) {
	log.Error(code, ` Error: `, message)
	resp := Error {
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Sorry, an unexpected error has occured.", http.StatusInternalServerError)
	}

	CustomErrorHandler = func(w http.ResponseWriter, status int, message string) {
		writeError(w, message, status)
	}
)
