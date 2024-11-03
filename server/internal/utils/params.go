package utils

import (
	"net/http"
	"server/internal/models"

	// "server/internal/models"

	"github.com/go-chi/chi"
)

func GetParam (req *http.Request, paramName string) string {
	param := chi.URLParam(req, paramName)
	return param
}

func GetPagination(req *http.Request) models.Paginators {
	pagination := req.Context().Value("pagination").(models.Paginators)
	return pagination
}
