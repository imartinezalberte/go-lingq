package course

import (
	"context"
)

type (
	PostCourseService interface {
		PostCourse(context.Context, CourseCommand) (Course, error)
	}

	postCourseService struct {
		r PostCourseRepoAction
	}

	Service struct {
		postCourseService PostCourseService
	}
)

func NewService(r Repo) Service {
	return Service{&postCourseService{r}}
}

func Execute[T interface{ ToCommand() any }](
	ctx context.Context,
	svc Service,
	cmd T,
) (any, error) {
	switch cmd := cmd.ToCommand().(type) {
	case CourseCommand:
		return svc.postCourseService.PostCourse(ctx, cmd)
	}
	return nil, nil
}

func (g *postCourseService) PostCourse(
	ctx context.Context,
	cmd CourseCommand,
) (Course, error) {
	return g.r.PostCourse(ctx, cmd)
}
