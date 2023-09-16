package get

import (
	"github.com/imartinezalberte/go-lingq/cmd/entities"
	"github.com/imartinezalberte/go-lingq/internal/search"
)

type SearchResources struct {
	Pagination entities.Pagination
	SortBy     entities.SortBy
	Type       entities.ResourceType
	Level      entities.ResourceLevel
	IsExternal bool
	IsPersonal bool
	Tags       []string
	Shelf      string
	Title      string
}

func (s SearchResources) ToCommand() (any, error) {
	sortBy, err := s.SortBy.ToDomain()
	if err != nil {
		return nil, err
	}

	resourceType, err := s.Type.ToDomain()
	if err != nil {
		return nil, err
	}

	level, err := s.Level.ToDomain()
	if err != nil {
		return nil, err
	}

	return search.ResourcesQuery{
		Pagination: s.Pagination.ToCommand(),
		SortBy:     sortBy,
		Type:       resourceType,
		Level:      level,
		IsExternal: s.IsExternal,
		IsPersonal: s.IsPersonal,
		Tags:       s.Tags,
		Shelf:      s.Shelf,
		TitleName:  s.Title,
	}, nil
}
