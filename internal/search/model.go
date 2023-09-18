package search

import "github.com/imartinezalberte/go-lingq/internal/entities"

type (
	SearchResource struct {
		ID           uint                    `json:"id"`
		Title        string                  `json:"title"`
		Type         entities.ResourceType   `json:"type"`
		Status       entities.ResourceStatus `json:"status"`
		Source       ResourceSource          `json:"source"`
		ImageURL     string                  `json:"imageUrl"`
		VideoURL     string                  `json:"videoUrl"`
		AudioURL     string                  `json:"audioUrl"`
		AudioPending bool                    `json:"audioPending"`
		SharedBy
		Description     string                 `json:"description"`
		CollectionID    uint                   `json:"collectionId"`
		CollectionTitle string                 `json:"collectionTitle"`
		NewWordsCount   uint                   `json:"newWordsCount"`
		Difficulty      float32                `json:"difficulty"`
		Level           entities.ResourceLevel `json:"level"`
		Date            entities.DateOnly      `json:"date"`
		Tags            []string               `json:"tags"`
		URL             string                 `json:"url"`
	}

	ResourceSource struct {
		// youtube, direct, netflix
		Type string `json:"type"`
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	SharedBy struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Role string `json:"role"`
	}
)
