package course

import (
	"context"

	"github.com/imartinezalberte/go-lingq/internal/entities"
)

type (
	PostCourseService interface {
		PostCourse(context.Context, CourseCommand) (Course, error)
	}

	postCourseService struct {
		r PostCourseRepoAction
	}

	GetCoursesService interface {
		GetCourses(context.Context, CourseQuery) (entities.Resource[Course], error)
	}

	getCoursesService struct {
		r GetCoursesRepoAction
	}

	GetCourseByIDService interface {
		GetCourse(context.Context, CourseQuery) (Course, error)
	}

	getCourseByIDService struct {
		r GetCourseByIDRepoAction
	}

	GetCoursesByIDService interface {
		GetCoursesByID(context.Context, CourseQuery) (CounterCourses, error)
	}

	getCoursesByIDService struct {
		r GetCoursesByIDRepoAction
	}

	Service struct {
		PostCourseService
		GetCoursesService
		GetCourseByIDService
		GetCoursesByIDService
	}
)

func NewService(r Repo) Service {
	return Service{
		&postCourseService{r},
		&getCoursesService{r},
		&getCourseByIDService{r},
		&getCoursesByIDService{r},
	}
}

func Execute[T interface{ ToCommand() any }](
	ctx context.Context,
	svc Service,
	cmd T,
) (any, error) {
	switch cmd := cmd.ToCommand().(type) {
	case CourseCommand:
		return svc.PostCourse(ctx, cmd)
	case CourseQuery:
		if len(cmd.IDs) > 1 {
			return svc.GetCoursesByID(ctx, cmd)
		} else if len(cmd.IDs) == 1 {
			return svc.GetCourse(ctx, cmd)
		}
		return svc.GetCourses(ctx, cmd)
	}
	return nil, nil
}

func (g *postCourseService) PostCourse(
	ctx context.Context,
	cmd CourseCommand,
) (Course, error) {
	return g.r.PostCourse(ctx, cmd)
}

func (g *getCoursesService) GetCourses(
	ctx context.Context,
	query CourseQuery,
) (entities.Resource[Course], error) {
	return g.r.GetCourses(ctx, query)
}

func (g *getCourseByIDService) GetCourse(ctx context.Context, query CourseQuery) (Course, error) {
	return g.r.GetCourseByID(ctx, query)
}

func (g *getCoursesByIDService) GetCoursesByID(
	ctx context.Context,
	query CourseQuery,
) (CounterCourses, error) {
	return g.r.GetCoursesByID(ctx, query)
}
