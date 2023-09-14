package language

import (
	lang "github.com/imartinezalberte/go-lingq/internal/language"
	"github.com/imartinezalberte/go-lingq/internal/utils"
)

const (
	CodeName      = "code"
	CodeShortName = "c"
	CodeUsage     = "Enter the code of the language for further use in the api that you are looking for"
	CodeDefault   = utils.Empty

	TitleName      = "name"
	TitleShortName = "n"
	TitleUsage     = "Enter the language name that you are looking for"
	TitleDefault   = utils.Empty

	SupportedName      = "supported"
	SupportedShortName = "s"
	SupportedUsage     = "Do you wanna display supported or unspported languages by lingq? By default is supported"
	SupportedDefault   = true

	IDName      = "id"
	IDShortName = "i"
	IDUsage     = "If you want to discover more information about a lingq language id, use this option"
	IDDefault   = 0

	KnownWordsName      = "known-words"
	KnownWordsShortName = "w"
	KnownWordsUsage     = "When you want to filter by >= amount of known words"
	KnownWordsDefault   = 0
)

type LanguageRequest struct {
	Code       string
	Title      string
	Supported  *bool
	ID         uint
	KnownWords uint
}

func (l LanguageRequest) ToCommand() any {
	if l.ID > 0 {
		return lang.LanguagesQuery{ID: l.ID}
	}

	return lang.LanguagesQuery{
		Code:       l.Code,
		Title:      l.Title,
		Supported:  l.Supported,
		KnownWords: l.KnownWords,
	}
}
