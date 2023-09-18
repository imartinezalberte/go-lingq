package entities

import (
	"errors"
	"strings"

	e "github.com/imartinezalberte/go-lingq/internal/entities"
	"github.com/spf13/cobra"
)

const (
	ResourceTypeName  = "resource-type"
	ResourceTypeUsage = "specify the type of resource that you want to look for: colletions or content"
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

func (r *ResourceType) Type() string {
	return "ResourceType"
}

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

func (r ResourceType) ToDomain() (t e.ResourceType, err error) {
	return t, t.Set(r.String())
}

func (r *ResourceType) Args(cmd *cobra.Command) {
	cmd.Flags().Var(r, ResourceTypeName, ResourceTypeUsage)
	cmd.RegisterFlagCompletionFunc(
		ResourceTypeName,
		func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
			return ResourceTypeValues[:], cobra.ShellCompDirectiveDefault
		},
	)
}
