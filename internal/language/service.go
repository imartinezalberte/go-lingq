package language

import (
	"context"

	"github.com/repeale/fp-go"
)

type (
	GetLanguagesService interface {
		GetLanguages(context.Context, LanguagesQuery) (Languages, error)
	}

	getLanguagesService struct {
		r GetLanguagesRepoAction
	}

	Service struct {
		getLanguagesService GetLanguagesService
	}
)

func NewService(r Repo) Service {
	return Service{&getLanguagesService{r}}
}

func Execute[T interface{ ToCommand() any }](
	ctx context.Context,
	svc Service,
	cmd T,
) (any, error) {
	switch cmd := cmd.ToCommand().(type) {
	case LanguagesQuery:
		return svc.getLanguagesService.GetLanguages(ctx, cmd)
	}
	return nil, nil
}

func (g *getLanguagesService) GetLanguages(
	ctx context.Context,
	query LanguagesQuery,
) (Languages, error) {
	languages, err := g.r.GetLanguages(ctx, query)
	if err != nil {
		return languages, err
	}

	return fp.Filter(query.Filter())(languages), err
}
