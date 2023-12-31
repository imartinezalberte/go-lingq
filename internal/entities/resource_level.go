package entities

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/imartinezalberte/go-lingq/internal/utils"
)

type ResourceLevel uint

const (
	_ ResourceLevel = iota
	FirstLevel
	SecondLevel
	ThirdLevel
	FourthLevel
	FifthLevel
	SixthLevel
)

var (
	Levels          = [...]string{"a1", "a2", "b1", "b2", "c1", "c2"}
	UnmarshalLevels = [...]string{
		"beginner 1",
		"beginner 2",
		"intermediate 1",
		"intermediate 2",
		"advanced 1",
		"advanced 2",
	}
)

func (r *ResourceLevel) Set(input string) error {
	if !r.Check(input) {
		return errors.New("unknown resource level")
	}
	return nil
}

func (r *ResourceLevel) Check(input string) bool {
	switch strings.TrimSpace(strings.ToLower(input)) {
	case UnmarshalLevels[0], Levels[0]:
		*r = FirstLevel
	case UnmarshalLevels[1], Levels[1]:
		*r = SecondLevel
	case UnmarshalLevels[2], Levels[2]:
		*r = ThirdLevel
	case UnmarshalLevels[3], Levels[3]:
		*r = FourthLevel
	case UnmarshalLevels[4], Levels[4]:
		*r = FifthLevel
	case UnmarshalLevels[5], Levels[5]:
		*r = SixthLevel
	default:
		return false
	}
	return true
}

func (r ResourceLevel) String() string {
	index := int(r) - 1
	if index >= len(Levels) {
		return "unknown"
	}

	return strings.ToUpper(Levels[index])
}

func (r ResourceLevel) EncodeValues(key string, v *url.Values) error {
	v.Add(key, strconv.Itoa(int(r)))
	return nil
}

func (r *ResourceLevel) UnmarshalJSON(input []byte) error {
	sanitized := strings.Trim(string(input), utils.Quote)
	if sanitized == "null" {
		return nil
	}

	return r.Set(sanitized)
}
