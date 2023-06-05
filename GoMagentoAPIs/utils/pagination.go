package utils

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Page     int
	PageSize int
}

func GetPagination(r *http.Request) (Pagination, error) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 10
	}

	return Pagination{
		Page:     page,
		PageSize: pageSize,
	}, nil
}