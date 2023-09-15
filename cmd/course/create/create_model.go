package create

import (
	"errors"
	"reflect"
	"strings"

	cour "github.com/imartinezalberte/go-lingq/internal/course"
	"github.com/imartinezalberte/go-lingq/internal/utils"
	"github.com/repeale/fp-go"
)

const (
	ImageName    = "image"
	ImageUsage   = "Specify here the filepath to the image that you want to upload"
	ImageDefault = utils.Empty

	TitleName    = "title"
	TitleUsage   = "Specify here the title of the course. It may contain spaces. Mandatory"
	TitleDefault = utils.Empty

	LanguageName    = "language"
	LanguageUsage   = "Specify the code of the language. It must be in two letter format. Mandatory"
	LanguageDefault = utils.Empty

	DescriptionName    = "description"
	DescriptionUsage   = "Specify here a description. It's optional. Max lenght must be	200 characters"
	DescriptionDefault = utils.Empty

	LevelName                  = "level"
	LevelUsage                 = "Specify here the level of your course."
	LevelDefault ResourceLevel = FirstLevel

	SourceURLName    = "source-url"
	SourceURLUsage   = "Specify the source URL of the course if appropiate"
	SourceURLDefault = utils.Empty

	TagsName  = "tags"
	TagsUsage = "Specify zero or more tags to classify the course. By default, empty"
)

type ResourceLevel uint

const (
	_ ResourceLevel = iota
	FirstLevel
	SecondLevel
	ThirdLevel
	FourthLevel
	FifthLevel
	SixthLevel
)

var Levels = [...]string{"A1", "A2", "B1", "B2", "C1", "C2"}

func (r *ResourceLevel) Type() string {
	return reflect.Uint.String()
}

func (r ResourceLevel) String() string {
	index := int(r) - 1
	if index >= len(Levels) || index < 0 {
		return "unknown"
	}

	return Levels[index]
}

func (r *ResourceLevel) Set(input string) error {
	if !r.Check(input) {
		return errors.New("resource level is not correct")
	}
	return nil
}

func (r *ResourceLevel) Check(input string) bool {
	switch strings.ToUpper(input) {
	case Levels[0]:
		*r = FirstLevel
	case Levels[1]:
		*r = SecondLevel
	case Levels[2]:
		*r = ThirdLevel
	case Levels[3]:
		*r = FourthLevel
	case Levels[4]:
		*r = FifthLevel
	case Levels[5]:
		*r = SixthLevel
	default:
		return false
	}
	return true
}

type CourseRequest struct {
	Image       string
	Title       string
	Language    string
	Description string
	Level       ResourceLevel
	SourceURL   string
	Tags        []string
}

func (c CourseRequest) ToCommand() any {
	return cour.CourseCommand{
		Image:       c.Image,
		Title:       c.Title,
		Language:    c.Language,
		Description: c.Description,
		Level:       uint(c.Level),
		SourceURL:   c.SourceURL,
		Tags: fp.Map(func(input string) cour.Tag {
			return cour.Tag(input)
		})(c.Tags),
	}
}
