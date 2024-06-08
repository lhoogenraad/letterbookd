package utils

import (
	"github.com/go-chi/chi"
	"net/http"
)

func GetParam (req *http.Request, paramName string) string {
	param := chi.URLParam(req, paramName)
	return param
}
