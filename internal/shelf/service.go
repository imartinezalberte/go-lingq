package shelf

import "context"

type (
	GetShelvesService interface {
		GetShelves(context.Context, ShelfQuery) (Shelves, error)
	}

	getShelvesService struct {
		r GetShelvesRepoAction
	}

	Service struct {
		getShelvesService GetShelvesService
	}
)

func NewService(r Repo) Service {
	return Service{&getShelvesService{r}}
}

func Execute[T interface{ ToCommand() any }](ctx context.Context, svc Service, cmd T) (any, error) {
	switch cmd := cmd.ToCommand().(type) {
	case ShelfQuery:
		return svc.getShelvesService.GetShelves(ctx, cmd)
	}
	return nil, nil
}

func (g *getShelvesService) GetShelves(ctx context.Context, query ShelfQuery) (Shelves, error) {
	return g.r.GetShelves(ctx, query)
}
