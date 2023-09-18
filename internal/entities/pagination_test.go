package entities_test

import (
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"

	"github.com/imartinezalberte/go-lingq/internal/entities"
)

var _ = Describe("Pagination", func() {
	var page, size uint

	BeforeEach(func() {
		page, size = 1, 10
	})

	When("everything goes find because the obj is created successfully", func() {
		It("works fine", func() {
			Ω(entities.NewPagination(page, size)).Should(MatchAllFields(Fields{
				"Page": Equal(page),
				"Size": Equal(size),
			}))
		})
	})

	DescribeTable(
		"page or/and size have not a correct value",
		func(input, output entities.Pagination) {
			Ω(entities.NewPagination(input.Page, input.Size)).Should(Equal(output))
		},
		Entry(
			"page has not a correct value",
			entities.Pagination{0, 10},
			entities.Pagination{entities.PageDefault, 10},
		),
		Entry(
			"pageSize has not a correct value",
			entities.Pagination{2, 0},
			entities.Pagination{2, entities.PageSizeDefault},
		),
		Entry(
			"page and size have not a correct value",
			entities.Pagination{},
			entities.Pagination{entities.PageDefault, entities.PageSizeDefault},
		),
	)

	When("querying, everything is right", func() {
		It("", func() {
			Ω(entities.NewPagination(page, size).ToQuery()).Should(MatchAllKeys(Keys{
				entities.PageQueryParamName:     Equal([]string{strconv.Itoa(int(page))}),
				entities.PageSizeQueryParamName: Equal([]string{strconv.Itoa(int(size))}),
			}))
		})
	})
})
