package entities

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/imartinezalberte/go-lingq/internal/utils"
)

type ResourceType uint

const (
	Course ResourceType = iota
	Lesson
)

var (
	ResourceTypeValues          = [...]string{"collection", "content"}
	ResourceTypeUnmarshalValues = [...]string{"courses", "lessons"}
)

func (r *ResourceType) Set(input string) error {
	if !r.Check(input) {
		return errors.New("unknown resource type")
	}
	return nil
}

func (r *ResourceType) Check(input string) bool {
	switch strings.TrimSpace(strings.ToLower(input)) {
	case ResourceTypeUnmarshalValues[0], ResourceTypeValues[0]:
		*r = Course
	case ResourceTypeUnmarshalValues[1], ResourceTypeValues[1]:
		*r = Lesson
	default:
		return false
	}
	return true
}

func (r ResourceType) String() string {
	t := int(r)
	if t >= len(ResourceTypeValues) {
		return "unknown"
	}

	return ResourceTypeValues[t]
}

func (r ResourceType) EncodeValues(key string, v *url.Values) error {
	v.Add(key, r.String())
	return nil
}

func (r *ResourceType) UnmarshalJSON(input []byte) error {
	return r.Set(strings.Trim(string(input), utils.Quote))
}

func (r ResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
