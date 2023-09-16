package search

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/entities"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

const (
	LanguageIDPathParam = "languageID"
	SearchEndpoint      = "/{" + LanguageIDPathParam + "}/search/"
)

type (
	GetResourcesRepoAction interface {
		GetResources(context.Context, ResourcesQuery) (entities.Resource[SearchResource], error)
	}

	Repo interface {
		GetResourcesRepoAction
	}

	repo struct {
		cl *resty.Client
	}
)

func NewRepo(cl *resty.Client) Repo {
	return &repo{cl}
}

func (r *repo) GetResources(
	ctx context.Context,
	query ResourcesQuery,
) (entities.Resource[SearchResource], error) {
	return rest.Exec[ResourcesQuery, rest.DummyAPIResponseErr, entities.Resource[SearchResource]](
		r.cl,
		ctx,
		query,
		SearchEndpoint,
	)
}
