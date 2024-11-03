package middleware

import (
	"context"
	"net/http"
	"strconv"
	"server/internal/resources"
)

var DEFAULT_PAGE int = 1
var MIN_PAGE int = 1

var DEFAULT_PAGE_SIZE int = 50
var MAX_PAGE_SIZE int = 500


func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		var pagination resources.Paginators
		pageStr := r.URL.Query().Get(string("page"))
		pageSizeStr := r.URL.Query().Get(string("pageSize"))

		pagination.Page, pagination.PageSize = parsePaginationValues(pageStr, pageSizeStr)

		ctx := context.WithValue(r.Context(), "pagination" ,pagination)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

/**
Converts given inputs from string to int.
Sets default values if errors occur
*/
func parsePaginationValues(pageString string, pageSizeString string) (int, int) {
	var page int = determinePage(pageString)
	var pageSize int = determinePageSize(pageSizeString)

	return page, pageSize
}


func determinePage(pageString string) int {
	var page int = DEFAULT_PAGE

	page, err := strconv.Atoi(pageString)
	if err != nil {page = DEFAULT_PAGE}

	if page < MIN_PAGE {page = DEFAULT_PAGE}

	return page
}

func determinePageSize(pageSizeString string) int {
	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {pageSize = DEFAULT_PAGE_SIZE}
	
	if pageSize > MAX_PAGE_SIZE {
		pageSize = MAX_PAGE_SIZE
	}

	return pageSize
}
