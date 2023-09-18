package shelf

import (
	"net/url"
	"strconv"

	"github.com/imartinezalberte/go-lingq/internal/entities"
	"github.com/imartinezalberte/go-lingq/internal/rest"
	"github.com/repeale/fp-go"
)

const LevelQueryParamName = "level"

type ShelfQuery struct {
	rest.GetDummyRequester
	entities.Pagination
	Levels   []uint
	Language string
}

func (s ShelfQuery) ToPathParameter() (map[string]string, error) {
	return map[string]string{LanguageIDPathParam: s.Language}, nil
}

func (s ShelfQuery) ToQuery() (url.Values, error) {
	paginationValues, err := s.Pagination.ToQuery()
	if err != nil {
		return paginationValues, err
	}

	if len(s.Levels) > 0 {
		paginationValues[LevelQueryParamName] = fp.Map(func(level uint) string {
			return strconv.Itoa(int(level))
		})(s.Levels)
	}

	return paginationValues, nil
}
