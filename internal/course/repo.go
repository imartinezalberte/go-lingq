package course

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/entities"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

const (
	LanguageIDPathParam = "languageID"
	CourseIDPathParam   = "courseID"

	CollectionQueryParam = "collection"

	PostCoursesEndpoint = "/{" + LanguageIDPathParam + "}/collections/"
	GetCoursesEndpoint  = "/{" + LanguageIDPathParam + "}/collections/{" + CourseIDPathParam + "}/"
)

type (
	PostCourseRepoAction interface {
		PostCourse(context.Context, CourseCommand) (Course, error)
	}

	GetCoursesRepoAction interface {
		GetCourses(context.Context, CourseQuery) (entities.Resource[Course], error)
	}

	GetCourseByIDRepoAction interface {
		GetCourseByID(context.Context, CourseQuery) (Course, error)
	}

	GetCoursesByIDRepoAction interface {
		GetCoursesByID(context.Context, CourseQuery) (CounterCourses, error)
	}

	Repo interface {
		PostCourseRepoAction
		GetCoursesRepoAction
		GetCourseByIDRepoAction
		GetCoursesByIDRepoAction
	}

	repo struct {
		cl *resty.Client
	}
)

func NewRepo(cl *resty.Client) Repo {
	return &repo{cl}
}

func (r *repo) PostCourse(ctx context.Context, cmd CourseCommand) (Course, error) {
	return rest.Exec[CourseCommand, rest.DummyAPIResponseErr, Course](
		r.cl,
		ctx,
		cmd,
		PostCoursesEndpoint,
	)
}

func (r *repo) GetCourses(
	ctx context.Context,
	query CourseQuery,
) (entities.Resource[Course], error) {
	return rest.Exec[CourseQuery, rest.DummyAPIResponseErr, entities.Resource[Course]](
		r.cl,
		ctx,
		query,
		GetCoursesEndpoint,
	)
}

func (r *repo) GetCourseByID(ctx context.Context, query CourseQuery) (Course, error) {
	return rest.Exec[CourseQuery, rest.DummyAPIResponseErr, Course](
		r.cl,
		ctx,
		query,
		GetCoursesEndpoint,
	)
}

func (r *repo) GetCoursesByID(ctx context.Context, query CourseQuery) (CounterCourses, error) {
	return rest.Exec[CourseQuery, rest.DummyAPIResponseErr, CounterCourses](
		r.cl,
		ctx,
		query,
		GetCoursesEndpoint,
	)
}
