package get

import (
	"github.com/imartinezalberte/go-lingq/internal/course"
	"github.com/imartinezalberte/go-lingq/internal/utils"
	"github.com/spf13/cobra"
)

const (
	LanguageName    = "language"
	LanguageUsage   = "Specify the language that you want to filter by"
	LanguageDefault = "en"

	TitleName    = "title"
	TitleUsage   = "Specify the words that you want the title contain"
	TitleDefault = utils.Empty

	IDName  = "id"
	IDUsage = "Specify the id to retrieve the exact match"
)

type CoursesRequest struct {
	IDs      []uint
	Title    string
	Language string
}

func (c CoursesRequest) ToCommand() any {
	return course.CourseQuery{
		IDs:      c.IDs,
		Title:    c.Title,
		Language: c.Language,
	}
}

func (c *CoursesRequest) Args(cmd *cobra.Command) {
	cmd.Flags().
		StringVar(&c.Title, TitleName, TitleDefault, TitleUsage)
	cmd.Flags().
		StringVar(&c.Language, LanguageName, LanguageDefault, LanguageUsage)
	cmd.Flags().
		UintSliceVar(&c.IDs, IDName, []uint{}, IDUsage)

	cmd.MarkFlagRequired(LanguageName)
}
