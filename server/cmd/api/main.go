package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"server/internal/handlers"
	"server/internal/tools"
	log "github.com/sirupsen/logrus"
)

func setupDatabase() error{
	var dbErr error
	tools.DB, dbErr = tools.NewDatabase()
	if dbErr != nil{
		log.Error(dbErr)
	}
	return dbErr
}

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()

	err := setupDatabase()
	if err != nil{
		os.Exit(1)
	}

	handlers.ApiHandler(r)
	log.Info("Starting the bombaclaat server")

	err = http.ListenAndServe("localhost:8080", r)

	if err != nil{
		log.Error(err)
	}
}
