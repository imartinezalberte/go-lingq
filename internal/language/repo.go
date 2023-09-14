package language

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

const GetLanguagesEndpoint = "languages/"

type (
	GetLanguagesRepoAction interface {
		GetLanguages(context.Context, LanguagesQuery) (Languages, error)
	}

	Repo interface {
		GetLanguagesRepoAction
	}

	repo struct {
		cl *resty.Client
	}
)

func NewRepo(cl *resty.Client) Repo {
	return &repo{cl}
}

func (r *repo) GetLanguages(ctx context.Context, query LanguagesQuery) (Languages, error) {
	return rest.Exec[LanguagesQuery, rest.DummyAPIResponseErr, Languages](
		r.cl,
		ctx,
		query,
		GetLanguagesEndpoint,
	)
}
