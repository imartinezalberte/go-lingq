package activity

import "github.com/imartinezalberte/go-lingq/internal/rest"

type ActivityQuery struct {
	rest.GetDummyRequester
	Language string
}

func (a ActivityQuery) ToPathParameter() (map[string]string, error) {
	return map[string]string{LanguageIDPathParam: a.Language}, nil
}
