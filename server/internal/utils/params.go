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

func GetPagination(req *http.Request) resources.Paginators {
	pagination := req.Context().Value("pagination").(resources.Paginators)
	return pagination
}
