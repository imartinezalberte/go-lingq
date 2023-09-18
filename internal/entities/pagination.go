package entities

import (
	"net/url"
	"strconv"
)

const (
	PageQueryParamName     = "page"
	PageSizeQueryParamName = "page_size"

	PageDefault     = 1
	PageSizeDefault = 20
)

type Pagination struct {
	Page uint `url:"page"`
	Size uint `url:"page_size"`
}

func NewPagination(page, size uint) Pagination {
	if page == 0 {
		page = PageDefault
	}

	if size < 1 {
		size = PageSizeDefault
	}

	return Pagination{page, size}
}

func (p Pagination) ToQuery() (url.Values, error) {
	return url.Values{
		PageQueryParamName:     []string{strconv.Itoa(int(p.Page))},
		PageSizeQueryParamName: []string{strconv.Itoa(int(p.Size))},
	}, nil
}
