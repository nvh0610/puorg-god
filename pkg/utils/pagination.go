package utils

import (
	"math"
	"net/url"
	"strconv"
)

func SetDefaultPagination(values url.Values) (int, int) {
	page, _ := strconv.Atoi(values.Get("page"))
	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(values.Get("limit"))
	if limit == 0 {
		limit = 10
	}

	return page, limit
}

func CalculatorTotalPage(total, limit int) int {
	return int(math.Ceil(float64(total) / float64(limit)))
}
