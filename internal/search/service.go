package search

import (
	"context"

	"github.com/imartinezalberte/go-lingq/internal/entities"
)

type (
	GetResourcesService interface {
		Execute(context.Context, ResourcesQuery) (entities.Resource[SearchResource], error)
	}

	getResourcesService struct {
		r GetResourcesRepoAction
	}

	Service struct {
		getResourcesService GetResourcesService
	}
)

func NewService(r Repo) Service {
	return Service{&getResourcesService{r}}
}

func Execute[T interface{ ToCommand() (any, error) }](
	ctx context.Context,
	svc Service,
	commander T,
) (any, error) {
	cmd, err := commander.ToCommand()
	if err != nil {
		return nil, err
	}

	switch cmd := cmd.(type) {
	case ResourcesQuery:
		return svc.getResourcesService.Execute(ctx, cmd)
	}

	return nil, nil
}

func (g *getResourcesService) Execute(
	ctx context.Context,
	query ResourcesQuery,
) (entities.Resource[SearchResource], error) {
	return g.r.GetResources(ctx, query)
}
