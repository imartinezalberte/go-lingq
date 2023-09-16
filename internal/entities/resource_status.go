package entities

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/imartinezalberte/go-lingq/internal/utils"
)

type ResourceStatus uint

const (
	PrivateResource ResourceStatus = iota
	SharedResource
)

var ResourceStatusValue = [...]string{"private", "shared"}

func (r ResourceStatus) String() string {
	t := int(r)
	if t >= len(ResourceStatusValue) {
		return "unknown"
	}

	return ResourceStatusValue[t]
}

func (r *ResourceStatus) Set(input string) error {
	if r.Check(input) {
		return errors.New("unknown resource status")
	}
	return nil
}

func (r *ResourceStatus) Check(input string) bool {
	switch strings.TrimSpace(strings.ToLower(input)) {
	case ResourceStatusValue[0]:
		*r = PrivateResource
	case ResourceStatusValue[1]:
		*r = SharedResource
	default:
		return false
	}
	return true
}

func (r ResourceStatus) EncodeValues(key string, v *url.Values) error {
	v.Add(key, r.String())
	return nil
}

func (r *ResourceStatus) UnmarshalJSON(input []byte) error {
	return r.Set(strings.Trim(string(input), utils.Quote))
}

func (r ResourceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
