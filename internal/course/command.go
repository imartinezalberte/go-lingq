package course

import (
	"net/url"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/rest"
	"github.com/imartinezalberte/go-lingq/internal/utils"
)

const ImageFileParamName string = "image"

type CourseCommand struct {
	rest.PostDummyRequester
	Image       string
	Title       string `example:"eventyr for barn"`
	Language    string `example:"no"`
	Description string `example:"Du vil lære å lese eventyr for barn"`
	Level       uint   `example:"2"`
	SourceURL   string `example:"https://www.barneforlaget.no/hør-så-mye-du-vil"`
	Tags        Tags
}

func (c CourseCommand) ToBody() (any, error) {
	if c.Image != utils.Empty {
		// We are not using simple json, we have to use multipart/form-data
		return nil, rest.ErrUnimplementedMethod
	}

	return AddCourse{
		Title:       c.Title,
		Language:    c.Language,
		Description: c.Description,
		Level:       c.Level,
		SourceURL:   c.SourceURL,
		Tags:        c.Tags,
	}, nil
}

func (c CourseCommand) ToPathParameter() (map[string]string, error) {
	return map[string]string{
		LanguageIDPathParam: c.Language,
	}, nil
}

func (c CourseCommand) After(req *resty.Request) (*resty.Request, error) {
	if c.Image == utils.Empty {
		return req, rest.ErrUnimplementedMethod
	}

	return req.
		SetFormDataFromValues(c.FormData()).
		SetFiles(c.Files()), nil
}

func (c CourseCommand) FormData() url.Values {
	return map[string][]string{
		"title":       {c.Title},
		"language":    {c.Language},
		"description": {c.Description},
		"level":       {strconv.Itoa(int(c.Level))},
		"sourceURL":   {c.SourceURL},
		"tags":        utils.SimilarString([]Tag(c.Tags)),
	}
}

func (c CourseCommand) Files() map[string]string {
	return map[string]string{ImageFileParamName: c.Image}
}
