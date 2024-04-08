package middleware

import (
	"errors"
	"net/http"

	"server/api"
	"server/internal/tools"

	log "github.com/sirupsen/logrus"
)

var UnauthorisedError = errors.New("Invalid username or token.")

func Authorisation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w httpResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnauthorisedError)
			api.RequestErrorHandler(w, UnauthorisedError)
			return
		}

		var db *tools.DatabaseInterface
		db, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}
		
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnauthorisedError)
			api.RequestErrorHandler(w, UnauthorisedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
