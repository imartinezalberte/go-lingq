package entities

import (
	e "github.com/imartinezalberte/go-lingq/internal/entities"
)

type SortBy uint

const (
	RecentlyOpenedSort SortBy = iota
	RecentlyImportedSort
	OldestSort
	NewestSort
	LessDifficultSort
	AlphabeticalSort
)

var SortByValues = [...]string{
	"recentlyOpened",
	"recentlyImported",
	"oldest",
	"newest",
	"lessDifficult",
	"alphabetical",
}

func (s SortBy) String() string {
	t := int(s)
	if t >= len(SortByValues) {
		return "unknown"
	}

	return SortByValues[t]
}

func (r SortBy) ToDomain() (sortBy e.SortBy, err error) {
	return sortBy, sortBy.Set(r.String())
}
