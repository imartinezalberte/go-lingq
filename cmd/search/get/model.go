package get

import (
	"github.com/imartinezalberte/go-lingq/cmd/entities"
	"github.com/imartinezalberte/go-lingq/internal/search"
	"github.com/imartinezalberte/go-lingq/internal/utils"
	"github.com/spf13/cobra"
)

const (
	IsExternalName    = "external"
	IsExternalUsage   = "the resource that you are looking for is external or not"
	IsExternalDefault = false

	IsPersonalName    = "personal"
	IsPersonalUsage   = "the resource that you are looking for is personal or not"
	IsPersonalDefault = false

	TagsName  = "tag"
	TagsUsage = "specify tags to query by"

	ShelfName    = "shelf"
	ShelfUsage   = "specify the shelf where you want to do the research"
	ShelfDefault = "my_lessons"

	TitleName    = "title"
	TitleUsage   = "specify the title of the resource"
	TitleDefault = utils.Empty

	LanguageName    = "language"
	LanguageUsage   = "Specify the code of the language. It must be in two letter format. Mandatory"
	LanguageDefault = utils.Empty
)

type SearchResources struct {
	Pagination entities.Pagination
	SortBy     entities.SortBy
	Type       entities.ResourceType
	Level      entities.ResourceLevel
	Language   string
	IsExternal bool
	IsPersonal bool
	Tags       []string
	Shelf      string
	Title      string
}

func (s SearchResources) ToCommand() (any, error) {
	sortBy, err := s.SortBy.ToDomain()
	if err != nil {
		return nil, err
	}

	resourceType, err := s.Type.ToDomain()
	if err != nil {
		return nil, err
	}

	level, err := s.Level.ToDomain()
	if err != nil {
		return nil, err
	}

	return search.ResourcesQuery{
		Pagination: s.Pagination.ToCommand(),
		SortBy:     sortBy,
		Type:       resourceType,
		Level:      level,
		IsExternal: s.IsExternal,
		IsPersonal: s.IsPersonal,
		Tags:       s.Tags,
		Shelf:      s.Shelf,
		TitleName:  s.Title,
		Language:   s.Language,
	}, nil
}

func (s *SearchResources) Args(cmd *cobra.Command) {
	s.Pagination.Args(cmd)
	s.SortBy.Args(cmd)
	s.Type.Args(cmd)
	s.Level.Args(cmd)

	cmd.Flags().BoolVar(&s.IsExternal, IsExternalName, IsExternalDefault, IsExternalUsage)
	cmd.Flags().BoolVar(&s.IsPersonal, IsPersonalName, IsPersonalDefault, IsPersonalUsage)
	cmd.Flags().StringSliceVar(&s.Tags, TagsName, []string{}, TagsUsage)
	cmd.Flags().StringVar(&s.Shelf, ShelfName, ShelfDefault, ShelfUsage)
	cmd.Flags().StringVar(&s.Title, TitleName, TitleDefault, TitleUsage)
	cmd.Flags().StringVar(&s.Language, LanguageName, LanguageDefault, LanguageUsage)

	cmd.MarkFlagRequired(LanguageName)
}
