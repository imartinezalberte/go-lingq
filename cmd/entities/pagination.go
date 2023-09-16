package entities

import (
	"fmt"

	"github.com/imartinezalberte/go-lingq/internal/pagination"
	"github.com/spf13/cobra"
)

const (
	PageName    = "page"
	PageUsage   = "set the page where to look for"
	PageDefault = 1

	SizeName    = "size"
	SizeUsage   = "set the max size that you want per size"
	SizeDefault = 20
)

type Pagination struct {
	Page, Size uint
}

func (p Pagination) String() string {
	return fmt.Sprintf("page: %d, size: %d", p.Page, p.Size)
}

func (p Pagination) ToCommand() pagination.Pagination {
	return pagination.NewPagination(p.Page, p.Size)
}

func (Pagination) Args(cmd *cobra.Command, target *Pagination) {
	cmd.Flags().UintVar(&target.Page, PageName, PageDefault, PageUsage)
	cmd.Flags().UintVar(&target.Size, SizeName, SizeDefault, SizeUsage)
}
