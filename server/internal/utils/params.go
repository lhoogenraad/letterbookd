package utils

import (
	"net/http"
	"server/internal/resources"


	"github.com/go-chi/chi"
)

func GetParam (req *http.Request, paramName string) string {
	param := chi.URLParam(req, paramName)
	return param
}

func GetUrlQuery (req *http.Request, queryParamName string) string {
	queryParam := req.URL.Query().Get(queryParamName)
	return queryParam
}

func GetPagination(req *http.Request) resources.Paginators {
	pagination := req.Context().Value("pagination").(resources.Paginators)
	return pagination
}
