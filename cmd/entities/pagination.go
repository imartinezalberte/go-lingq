package entities

import (
	"fmt"

	e "github.com/imartinezalberte/go-lingq/internal/entities"
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

func (p Pagination) ToCommand() e.Pagination {
	return e.NewPagination(p.Page, p.Size)
}

func (p *Pagination) Args(cmd *cobra.Command) {
	cmd.Flags().UintVar(&p.Page, PageName, PageDefault, PageUsage)
	cmd.Flags().UintVar(&p.Size, SizeName, SizeDefault, SizeUsage)
}
