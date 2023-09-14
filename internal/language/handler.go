package language

const (
	CodeName      = "code"
	CodeShortName = "c"
	CodeUsage     = "Enter the code of the language for further use in the api that you are looking for"

	TitleName      = "name"
	TitleShortName = "n"
	TitleUsage     = "Enter the language name that you are looking for"

	SupportedName      = "supported"
	SupportedShortName = "s"
	SupportedUsage     = "Do you wanna display supported or unspported languages by lingq? By default is supported"

	IDName      = "id"
	IDShortName = "i"
	IDUsage     = "If you want to discover more information about a lingq language id, use this option"
)

type LanguageRequest struct {
	Code      string
	Title     string
	Supported *bool
	ID        int
}

func (l LanguageRequest) ToCommand() any {
	if l.ID > 0 {
		return LanguagesQuery{ID: l.ID}
	}

	return LanguagesQuery{
		Code:      l.Code,
		Title:     l.Title,
		Supported: l.Supported,
	}
}
