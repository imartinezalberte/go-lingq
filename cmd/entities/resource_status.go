package entities

import (
	"errors"
	"strings"

	e "github.com/imartinezalberte/go-lingq/internal/entities"
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

func (r ResourceStatus) ToDomain() (status e.ResourceStatus, err error) {
	return status, status.Set(r.String())
}
