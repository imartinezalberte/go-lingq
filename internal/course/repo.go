package course

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

const (
	LanguageIDPathParam = "languageID"
	PostCoursesEndpoint = "/{" + LanguageIDPathParam + "}/collections/"
	GetCoursesEndpoint  = "/{" + LanguageIDPathParam + "}/collections/"
)

type (
	PostCourseRepoAction interface {
		PostCourse(context.Context, CourseCommand) (Course, error)
	}

	GetCoursesRepoAction interface {
		GetCourses(context.Context, CourseQuery) (Courses, error)
	}

	Repo interface {
		PostCourseRepoAction
		GetCoursesRepoAction
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

func (r *repo) GetCourses(ctx context.Context, query CourseQuery) (Courses, error) {
	return rest.Exec[CourseQuery, rest.DummyAPIResponseErr, Courses](
		r.cl,
		ctx,
		query,
		GetCoursesEndpoint,
	)
}
