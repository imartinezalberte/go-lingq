package activity

import "context"

type (
	GetActivityService interface {
		GetActivity(context.Context, ActivityQuery) (Activity, error)
	}

	getActivityService struct {
		r GetActivityRepoAction
	}

	Service struct {
		GetActivityService
	}
)

func NewService(r Repo) Service {
	return Service{&getActivityService{r}}
}

func Execute[T interface{ ToCommand() any }](
	ctx context.Context,
	svc Service,
	cmd T,
) (any, error) {
	switch cmd := cmd.ToCommand().(type) {
	case ActivityQuery:
		return svc.GetActivity(ctx, cmd)
	}
	return nil, nil
}

func (g *getActivityService) GetActivity(
	ctx context.Context,
	query ActivityQuery,
) (Activity, error) {
	return g.r.GetActivity(ctx, query)
}
