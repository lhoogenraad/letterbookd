package middleware

import (
	"errors"
	"context"
	"strings"
	"net/http"

	"server/api"
	"server/internal/utils"
	"github.com/dgrijalva/jwt-go"
	
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
		var tokenHeader string = r.Header.Get("Authorization")
		headerSplit := strings.Split(tokenHeader, "Bearer ")
		if len(headerSplit) != 2{
			log.Error("Malformed token: " + tokenHeader)
			api.CustomErrorHandler(w, 401, "Malformed token detected. Token must be in format 'Bearer <token>'")
			return
		}

		var token string = headerSplit[1]
		var err error

		claims, err := utils.VerifyToken(token)
		if err != nil {
			log.Error(err)
			api.RequestErrorHandler(w, UnauthorisedError)
			return
		}

		// Add claims map to request context for further handlers
		claimsContext := getTokenClaimsContext(claims, r)
		next.ServeHTTP(w, r.WithContext(claimsContext))
	})
}

func getTokenClaimsContext(claims jwt.MapClaims, r *http.Request) context.Context {
	ctx := context.WithValue(r.Context(), "claims", claims)
	return ctx
}
