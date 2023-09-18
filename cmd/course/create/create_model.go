package create

import (
	"github.com/imartinezalberte/go-lingq/cmd/entities"
	cour "github.com/imartinezalberte/go-lingq/internal/course"
	"github.com/imartinezalberte/go-lingq/internal/utils"
	"github.com/repeale/fp-go"
	"github.com/spf13/cobra"
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

	SourceURLName    = "source-url"
	SourceURLUsage   = "Specify the source URL of the course if appropiate"
	SourceURLDefault = utils.Empty

	TagsName  = "tags"
	TagsUsage = "Specify zero or more tags to classify the course. By default, empty"
)

type CourseRequest struct {
	Image       string
	Title       string
	Language    string
	Description string
	Level       entities.ResourceLevel
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

func (c *CourseRequest) Args(cmd *cobra.Command) {
	c.Level.Args(cmd)

	cmd.Flags().
		StringVar(&c.Title, TitleName, TitleDefault, TitleUsage)
	cmd.Flags().
		StringVar(&c.Language, LanguageName, LanguageDefault, LanguageUsage)
	cmd.Flags().
		StringVar(&c.Description, DescriptionName, DescriptionDefault, DescriptionUsage)
	cmd.Flags().
		StringVar(&c.SourceURL, SourceURLName, SourceURLDefault, SourceURLUsage)
	cmd.Flags().
		StringSliceVar(&c.Tags, TagsName, []string{}, TagsUsage)
	cmd.Flags().StringVar(&c.Image, ImageName, ImageDefault, ImageUsage)

	cmd.MarkFlagFilename(ImageName)

	cmd.MarkFlagRequired(TitleName)
	cmd.MarkFlagRequired(LanguageName)
}
