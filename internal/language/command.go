package language

import (
	"strings"

	"github.com/imartinezalberte/go-lingq/internal/rest"
	"github.com/imartinezalberte/go-lingq/internal/utils"
)

type LanguagesQuery struct {
	rest.GetDummyRequester
	Code, Title string
	Supported   *bool
	ID          uint
	KnownWords  uint
}

func (lq LanguagesQuery) Filter() utils.Predicate[Language] {
	if lq.ID > 0 {
		return func(l Language) bool {
			return l.ID == lq.ID
		}
	}

	conditions := make([]utils.Predicate[Language], 0, 3)

	if code := strings.ToLower(strings.TrimSpace(lq.Code)); code != utils.Empty {
		conditions = append(conditions, func(l Language) bool {
			return strings.Contains(strings.ToLower(l.Code), code)
		})
	}

	if title := strings.ToLower(strings.TrimSpace(lq.Title)); title != utils.Empty {
		conditions = append(conditions, func(l Language) bool {
			return strings.Contains(strings.ToLower(l.Title), title)
		})
	}

	if lq.KnownWords > 0 {
		conditions = append(conditions, func(l Language) bool {
			return l.KnownWords >= lq.KnownWords
		})
	}

	if lq.Supported != nil {
		conditions = append(conditions, func(l Language) bool {
			return *lq.Supported == l.Supported
		})
	}

	return func(l Language) bool {
		for _, condition := range conditions {
			if !condition(l) {
				return false
			}
		}

		return true
	}
}
