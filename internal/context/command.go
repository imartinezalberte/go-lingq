package language

import (
	"strings"

	"github.com/imartinezalberte/go-lingq/internal/language"
	"github.com/imartinezalberte/go-lingq/internal/rest"
	"github.com/imartinezalberte/go-lingq/internal/utils"
)

type ContextsQuery struct {
	rest.GetDummyRequester
	Identifier uint
	StreakDays uint
	Language   language.LanguagesQuery
	Intense    string // Casual, Steady, Intense, Insane
}

func (cq ContextsQuery) Filter() utils.Predicate[Context] {
	if cq.Identifier > 0 {
		return func(c Context) bool {
			return cq.Identifier == c.Identifier
		}
	}

	conditions := make([]utils.Predicate[Context], 0, 3)

	if cq.StreakDays > 0 {
		conditions = append(conditions, func(c Context) bool {
			return c.StreakDays >= cq.StreakDays
		})
	}

	if intense := strings.ToLower(strings.TrimSpace(cq.Intense)); intense != utils.Empty {
		conditions = append(conditions, func(c Context) bool {
			return strings.Contains(strings.ToLower(c.Intense), intense)
		})
	}

	languageFilter := cq.Language.Filter()
	conditions = append(conditions, func(c Context) bool {
		return languageFilter(c.Language)
	})

	return func(c Context) bool {
		for _, condition := range conditions {
			if !condition(c) {
				return false
			}
		}

		return true
	}
}
