package get

import (
	"github.com/imartinezalberte/go-lingq/cmd/entities"
	"github.com/imartinezalberte/go-lingq/internal/shelf"
	"github.com/imartinezalberte/go-lingq/internal/utils"
	"github.com/repeale/fp-go"
	"github.com/spf13/cobra"
)

const (
	LanguageName    = "language"
	LanguageUsage   = "Specify the code of the language. It must be in two letter format. Mandatory"
	LanguageDefault = utils.Empty

	LevelsName  = "levels"
	LevelsUsage = "Specify the levels in which you want to query shelves. Default is A1"
)

var LevelsDefault = []entities.ResourceLevel{entities.FirstLevel}

type ShelfRequest struct {
	Pagination entities.Pagination
	Levels     entities.ResourcesLevel
	Language   string
}

func (s ShelfRequest) String() string {
	return s.Pagination.String() + utils.Space + s.Levels.String() + utils.Space + s.Language
}

func (s ShelfRequest) ToCommand() any {
	return shelf.ShelfQuery{
		Pagination: s.Pagination.ToCommand(),
		Language:   s.Language,
		Levels: fp.Map(func(level entities.ResourceLevel) uint {
			return uint(level)
		})(s.Levels.InnerType().ToArr()),
	}
}

func (s *ShelfRequest) Args(cmd *cobra.Command) {
	cmd.Flags().
		StringVar(&s.Language, LanguageName, LanguageDefault, LanguageUsage)
	cmd.Flags().Var(&s.Levels, entities.LevelName, entities.LevelUsage)

	s.Pagination.Args(cmd)

	cmd.MarkFlagRequired(LanguageName)
}
