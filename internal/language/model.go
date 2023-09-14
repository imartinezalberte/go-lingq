package language

type (
	Languages []Language
	Language  struct {
		ID         uint   `json:"id"         example:"53"`
		URL        string `json:"url"        example:"https://www.lingq.com/api/v2/languages/{language_id}/"`
		Code       string `json:"code"       example:"mk"`
		Title      string `json:"title"      example:"Macedonian"`
		Supported  bool   `json:"supported"  example:"false"`
		KnownWords uint   `json:"knownWords" example:"0"`
	}
)
