package shelf

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

const (
	LanguageIDPathParam = "languageID"
	GetShelvesEndpoint  = "{" + LanguageIDPathParam + "}/shelves/"
)

type (
	GetShelvesRepoAction interface {
		GetShelves(context.Context, ShelfQuery) (Shelves, error)
	}

	Repo interface {
		GetShelvesRepoAction
	}

	repo struct {
		cl *resty.Client
	}
)

func NewRepo(cl *resty.Client) Repo {
	return &repo{cl}
}

func (r *repo) GetShelves(ctx context.Context, query ShelfQuery) (Shelves, error) {
	return rest.Exec[ShelfQuery, rest.DummyAPIResponseErr, Shelves](
		r.cl,
		ctx,
		query,
		GetShelvesEndpoint,
	)
}
