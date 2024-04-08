package middleware

import (
	"errors"
	"net/http"

	"server/api"
	"server/internal/utils"
	
	log "github.com/sirupsen/logrus"
)


type TokenClaims struct {
	userId		int		`json:userid`
	username	string	`json:username`
	email		string	`json:email`
}

var secretToken = []byte("secret!")

var UnauthorisedError = errors.New("Invalid username or token.")

func Authorisation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string = r.Header.Get("Authorization")
		var err error

		_, err = utils.VerifyToken(token)
		if err != nil {
			log.Error(err)
			api.RequestErrorHandler(w, UnauthorisedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
