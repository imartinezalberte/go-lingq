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

	GetCoursesService interface {
		GetCourses(context.Context, CourseQuery) (Courses, error)
	}

	getCoursesService struct {
		r GetCoursesRepoAction
	}

	Service struct {
		postCourseService PostCourseService
		getCoursesService GetCoursesService
	}
)

func NewService(r Repo) Service {
	return Service{&postCourseService{r}, &getCoursesService{r}}
}

func Execute[T interface{ ToCommand() any }](
	ctx context.Context,
	svc Service,
	cmd T,
) (any, error) {
	switch cmd := cmd.ToCommand().(type) {
	case CourseCommand:
		return svc.postCourseService.PostCourse(ctx, cmd)
	case CourseQuery:
		return svc.getCoursesService.GetCourses(ctx, cmd)
	}
	return nil, nil
}

func (g *postCourseService) PostCourse(
	ctx context.Context,
	cmd CourseCommand,
) (Course, error) {
	return g.r.PostCourse(ctx, cmd)
}

func (g *getCoursesService) GetCourses(ctx context.Context, query CourseQuery) (Courses, error) {
	return Courses{}, nil
}
