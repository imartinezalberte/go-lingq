package entities

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
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

func (s *SortBy) Set(input string) error {
	if !s.Check(input) {
		return errors.New("unkown sort type")
	}
	return nil
}

func (s *SortBy) Check(input string) bool {
	switch strings.TrimSpace(input) {
	case SortByValues[0], "0":
		*s = RecentlyOpenedSort
	case SortByValues[1], "1":
		*s = RecentlyImportedSort
	case SortByValues[2], "2":
		*s = OldestSort
	case SortByValues[3], "3":
		*s = NewestSort
	case SortByValues[4], "4":
		*s = LessDifficultSort
	case SortByValues[4], "5":
		*s = AlphabeticalSort
	default:
		return false
	}
	return true
}

func (s SortBy) String() string {
	t := int(s)
	if t >= len(SortByValues) {
		return "unknown"
	}

	return SortByValues[t]
}

func (s SortBy) EncodeValues(key string, v *url.Values) error {
	v.Add(key, s.String())
	return nil
}

func (s SortBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
