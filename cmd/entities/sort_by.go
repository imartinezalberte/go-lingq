package entities

import (
	"errors"
	"strings"

	e "github.com/imartinezalberte/go-lingq/internal/entities"
	"github.com/spf13/cobra"
)

const (
	SortByName  = "sort-by"
	SortByUsage = "sort results by availables options"
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

func (s *SortBy) Type() string {
	return "SortBy"
}

func (s *SortBy) Set(input string) error {
	if !s.Check(input) {
		return errors.New("unknown sort type")
	}
	return nil
}

func (s *SortBy) Check(input string) bool {
	switch strings.TrimSpace(input) {
	case SortByValues[0]:
		*s = RecentlyOpenedSort
	case SortByValues[1]:
		*s = RecentlyImportedSort
	case SortByValues[2]:
		*s = OldestSort
	case SortByValues[3]:
		*s = NewestSort
	case SortByValues[4]:
		*s = LessDifficultSort
	case SortByValues[5]:
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

func (r SortBy) ToDomain() (sortBy e.SortBy, err error) {
	return sortBy, sortBy.Set(r.String())
}

func (r *SortBy) Args(cmd *cobra.Command) {
	cmd.Flags().Var(r, SortByName, SortByUsage)
	cmd.RegisterFlagCompletionFunc(
		SortByName,
		func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
			return SortByValues[:], cobra.ShellCompDirectiveDefault
		},
	)
}
