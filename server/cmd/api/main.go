package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"server/internal/handlers"
	"server/internal/tools"
	log "github.com/sirupsen/logrus"
)


func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	dbErr := tools.NewDatabase()
	if dbErr != nil{
		log.Error(dbErr)
	}
	handlers.ApiHandler(r)
	log.Info("Starting the bombaclaat server")

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil{
		log.Error(err)
	}
}
