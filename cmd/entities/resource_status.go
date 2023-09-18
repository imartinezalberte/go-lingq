package entities

import (
	"errors"
	"strings"

	e "github.com/imartinezalberte/go-lingq/internal/entities"
	"github.com/spf13/cobra"
)

const (
	ResourceStatusName  = "resource-status"
	ResourceStatusUsage = "specify the status of the resource: private or shared"
)

type ResourceStatus uint

const (
	PrivateResource ResourceStatus = iota
	SharedResource
)

var ResourceStatusValue = [...]string{"private", "shared"}

func (r *ResourceStatus) Type() string {
	return "ResourceStatus"
}

func (r ResourceStatus) String() string {
	t := int(r)
	if t >= len(ResourceStatusValue) {
		return "unknown"
	}

	return ResourceStatusValue[t]
}

func (r *ResourceStatus) Set(input string) error {
	if !r.Check(input) {
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

func (r *ResourceStatus) Args(cmd *cobra.Command) {
	cmd.Flags().Var(r, ResourceStatusName, ResourceStatusUsage)
	cmd.RegisterFlagCompletionFunc(
		ResourceStatusName,
		func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
			return ResourceStatusValue[:], cobra.ShellCompDirectiveDefault
		},
	)
}
