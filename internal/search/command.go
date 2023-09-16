package search

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/imartinezalberte/go-lingq/internal/entities"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

type ResourcesQuery struct {
	rest.GetDummyRequester `url:"-"`
	Pagination             entities.Pagination
	SortBy                 entities.SortBy        `url:"sortBy"`
	Type                   entities.ResourceType  `url:"type"`
	Level                  entities.ResourceLevel `url:"level"`
	IsExternal             bool                   `url:"isExternal"`
	IsPersonal             bool                   `url:"isPersonal"`
	Tags                   []string               `url:"tags"`
	Shelf                  string                 `url:"shelf"`
	TitleName              string                 `url:"q"`
}

func (r ResourcesQuery) ToQuery() (url.Values, error) {
	return query.Values(r)
}
