package course

import (
	cour "github.com/imartinezalberte/go-lingq/internal/course"
	"github.com/imartinezalberte/go-lingq/internal/utils"
	"github.com/repeale/fp-go"
)

const (
	ImageName    = "image"
	ImageUsage   = "Specify here the filepath to the image that you want to upload"
	ImageDefault = utils.Empty

	TitleName  = "title"
	TitleUsage = "Specify here the title of the course. It may contain spaces. Mandatory"

	LanguageName  = "language"
	LanguageUsage = "Specify the code of the language. It must be in two letter format. Mandatory"

	DescriptionName    = "description"
	DescriptionUsage   = "Specify here a description. It's optional. Max lenght must be	200 characters"
	DescriptionDefault = utils.Empty

	LevelName    = "level"
	LevelUsage   = "Specify here the level of your course."
	LevelDefault = 1

	SourceURLName  = "source-url"
	SourceURLUsage = "Specify the source URL of the course if appropiate"

	TagsName  = "tags"
	TagsUsage = "Specify zero or more tags to classify the course. By default, empty"
)

type CourseRequest struct {
	Image       string
	Title       string
	Language    string
	Description string
	Level       uint
	SourceURL   string
	Tags        []string
}

func (c CourseRequest) ToCommand() any {
	return cour.CourseCommand{
		Image:       c.Image,
		Title:       c.Title,
		Language:    c.Language,
		Description: c.Description,
		Level:       c.Level,
		SourceURL:   c.SourceURL,
		Tags: fp.Map(func(input string) cour.Tag {
			return cour.Tag(input)
		})(c.Tags),
	}
}
