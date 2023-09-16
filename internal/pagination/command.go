package pagination

import (
	"net/url"
	"strconv"
)

const (
	PageQueryParamName     = "page"
	PageSizeQueryParamName = "page_size"
)

type Pagination struct {
	Page uint
	Size uint
}

func NewPagination(page, size uint) Pagination {
	if page == 0 {
		page = 1
	}

	if size < 1 {
		size = 20
	}

	return Pagination{page, size}
}

func (p Pagination) ToQuery() (url.Values, error) {
	return url.Values{
		PageQueryParamName:     []string{strconv.Itoa(int(p.Page))},
		PageSizeQueryParamName: []string{strconv.Itoa(int(p.Size))},
	}, nil
}
