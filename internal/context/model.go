package language

import "github.com/imartinezalberte/go-lingq/internal/language"

type (
	ContextRes struct {
		Count   uint     `json:"count"`
		Results Contexts `json:"results"`
	}

	Contexts []Context
	Context  struct {
		Identifier uint              `json:"pk"          example:"5857878"`
		URL        string            `json:"url"         example:"https://www.lingq.com/api/v2/contexts/{context_identifier}/"`
		Language   language.Language `json:"language"`
		Intense    string            `json:"intense"     example:"insane"`
		StreakDays uint              `json:"streak_days" example:"20"`
		AppleLevel int               `json:"apple_level" example:"5"`
	}
)
