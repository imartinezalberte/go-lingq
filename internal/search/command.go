package search

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/imartinezalberte/go-lingq/internal/entities"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

type ResourcesQuery struct {
	rest.GetDummyRequester `url:"-"`
	entities.Pagination
	SortBy     entities.SortBy        `url:"sortBy"`
	Type       entities.ResourceType  `url:"type"`
	Level      entities.ResourceLevel `url:"level"`
	Language   string                 `url:"-"`
	IsExternal bool                   `url:"isExternal,omitempty"`
	IsPersonal bool                   `url:"isPersonal,omitempty"`
	Tags       []string               `url:"tags"`
	Shelf      string                 `url:"shelf"`
	TitleName  string                 `url:"q"`
}

func (r ResourcesQuery) ToPathParameter() (map[string]string, error) {
	return map[string]string{LanguageIDPathParam: r.Language}, nil
}

func (r ResourcesQuery) ToQuery() (url.Values, error) {
	return query.Values(r)
}
