package handlers

import (
	"encoding/json"
	"net/http"

	"server/api"
	"server/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}

	var decoder *schema.Decoder = schema.NewDecoder()

	var err error
	
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		// Change this to http.StatusOk at some point!!
		Code: 200,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	log.Info("Got balance for %s successfully.", params.Username)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
