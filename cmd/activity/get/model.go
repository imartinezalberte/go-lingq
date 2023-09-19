package get

import (
	"github.com/imartinezalberte/go-lingq/internal/activity"
	"github.com/imartinezalberte/go-lingq/internal/utils"
	"github.com/spf13/cobra"
)

const (
	LanguageName  = "language"
	LanguageUsage = "specify the language from which you want to get activity information"
)

type ActivityRequest struct {
	Language string
}

func (a ActivityRequest) ToCommand() any {
	return activity.ActivityQuery{Language: a.Language}
}

func (a *ActivityRequest) Args(cmd *cobra.Command) {
	cmd.Flags().StringVar(&a.Language, LanguageName, utils.Empty, LanguageUsage)

	cmd.MarkFlagRequired(LanguageName)
}
