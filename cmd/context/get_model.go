package context

import (
	"github.com/imartinezalberte/go-lingq/cmd/language"
	con "github.com/imartinezalberte/go-lingq/internal/context"
	lang "github.com/imartinezalberte/go-lingq/internal/language"
	"github.com/imartinezalberte/go-lingq/internal/utils"
)

const (
	IdentifierName    = "context-id"
	IdentifierUsage   = "Identifier of the context. Use it when you know exactly which one is"
	IdentifierDefault = 0

	StreakDaysName    = "streak-days"
	StreakDaysUsage   = "What amount of streak days you wanna look for?"
	StreakDaysDefault = 0

	IntenseName    = "intense"
	IntenseUsage   = "What type of intensity are you looking for? Casual, Steady, Intense or Insane (If you don't select none of those, empty is used)"
	IntenseDefault = utils.Empty
)

type ContextRequest struct {
	Identifier uint
	StreakDays uint
	Language   language.LanguageRequest
	Intense    string // Casual, Steady, Intense, Insane
}

func (c ContextRequest) ToCommand() any {
	if c.Identifier > 0 {
		return con.ContextsQuery{Identifier: c.Identifier}
	}

	return con.ContextsQuery{
		StreakDays: c.StreakDays,
		Language:   c.Language.ToCommand().(lang.LanguagesQuery),
		Intense:    c.Intense,
	}
}
