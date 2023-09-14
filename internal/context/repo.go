package language

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

const GetContextsEndpoint = "contexts/"

type (
	GetContextsRepoAction interface {
		GetContexts(context.Context, ContextsQuery) (ContextRes, error)
	}

	Repo interface {
		GetContextsRepoAction
	}

	repo struct {
		cl *resty.Client
	}
)

func NewRepo(cl *resty.Client) Repo {
	return &repo{cl}
}

func (r *repo) GetContexts(ctx context.Context, query ContextsQuery) (ContextRes, error) {
	return rest.Exec[ContextsQuery, rest.DummyAPIResponseErr, ContextRes](
		r.cl,
		ctx,
		query,
		GetContextsEndpoint,
	)
}
