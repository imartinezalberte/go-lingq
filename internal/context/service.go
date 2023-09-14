package language

import (
	"context"

	"github.com/repeale/fp-go"
)

type (
	GetContextsService interface {
		GetContexts(context.Context, ContextsQuery) (Contexts, error)
	}

	getContextsService struct {
		r GetContextsRepoAction
	}

	Service struct {
		getContextsService GetContextsService
	}
)

func NewService(r Repo) Service {
	return Service{&getContextsService{r}}
}

func Execute[T interface{ ToCommand() any }](
	ctx context.Context,
	svc Service,
	cmd T,
) (any, error) {
	switch cmd := cmd.ToCommand().(type) {
	case ContextsQuery:
		return svc.getContextsService.GetContexts(ctx, cmd)
	}
	return nil, nil
}

func (g *getContextsService) GetContexts(
	ctx context.Context,
	query ContextsQuery,
) (Contexts, error) {
	contexts, err := g.r.GetContexts(ctx, query)
	if err != nil {
		return Contexts{}, err
	}

	return fp.Filter(query.Filter())(contexts.Results), err
}
