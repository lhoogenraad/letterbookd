package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/lhoogenraad/letterbookd/internal/handlers"
	"github.com/lhoogenraad/letterbookd/internal/tools"
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

	port := os.Getenv("PORT")
	addr := "[::]:" + port
	log.Info("Attempting to listen on ", addr)
	err = http.ListenAndServe(addr, r)

	if err != nil{
		log.Error(err)
	}
}
