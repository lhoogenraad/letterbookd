package utils

import (
	"fmt"
	"net/http"
	// "server/internal/models"

	"github.com/go-chi/chi"
)

func GetParam (req *http.Request, paramName string) string {
	param := chi.URLParam(req, paramName)
	return param
}

func GetPagination(req *http.Request) {
	d := req.Context().Value("pagination")
	fmt.Println("D", d)
}
