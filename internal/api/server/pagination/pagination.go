package pagination

import (
	"net/http"
	"strconv"
)

func Paginate(r *http.Request) (int, int, int) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 250
	}

	offset := (page - 1) * limit

	return limit, offset, page
}
