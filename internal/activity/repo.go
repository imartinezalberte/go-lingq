package activity

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

const (
	LanguageIDPathParam = "languageID"
	GetActivityEndpoint = "/{" + LanguageIDPathParam + "}/study-stats/"
)

type (
	GetActivityRepoAction interface {
		GetActivity(context.Context, ActivityQuery) (Activity, error)
	}

	Repo interface {
		GetActivityRepoAction
	}

	repo struct {
		cl *resty.Client
	}
)

func NewRepo(cl *resty.Client) Repo {
	return &repo{cl}
}

func (r *repo) GetActivity(ctx context.Context, query ActivityQuery) (Activity, error) {
	return rest.Exec[ActivityQuery, rest.DummyAPIResponseErr, Activity](
		r.cl,
		ctx,
		query,
		GetActivityEndpoint,
	)
}
