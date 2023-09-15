package course

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

const (
	LanguageIDPathParam = "languageID"
	PostCoursesEndpoint = "/{" + LanguageIDPathParam + "}/collections/"
)

type (
	PostCourseRepoAction interface {
		PostCourse(context.Context, CourseCommand) (Course, error)
	}

	Repo interface {
		PostCourseRepoAction
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
