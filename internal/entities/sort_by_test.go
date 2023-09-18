package entities_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go-lingq/internal/entities"
)

var _ = Describe("SortBy", func() {
	var sortBy entities.SortBy

	When("Input is not correct", func() {
		It("should return error", func() {
			Ω(sortBy.Set("incorrect input")).Should(MatchError("unknown sort type"))
		})
	})

	When("Content is incorrect", func() {
		BeforeEach(func() {
			sortBy = entities.SortBy(len(entities.SortByValues))
		})

		It("should return \"unknown\" string", func() {
			Ω(sortBy.String()).Should(Equal("unknown"))
		})
	})

	DescribeTable(
		"String should return the right type",
		func(input entities.SortBy, output string) {
			Ω(input.String()).Should(Equal(output))
		},
		Entry("Recently opened", entities.RecentlyOpenedSort, entities.SortByValues[0]),
		Entry("recently imported", entities.RecentlyImportedSort, entities.SortByValues[1]),
		Entry("oldest", entities.OldestSort, entities.SortByValues[2]),
		Entry("newest", entities.NewestSort, entities.SortByValues[3]),
		Entry("less difficult", entities.LessDifficultSort, entities.SortByValues[4]),
		Entry("alphabetical", entities.AlphabeticalSort, entities.SortByValues[5]),
	)

	DescribeTable(
		"Set works with every possible option",
		func(output entities.SortBy, input string) {
			Ω(sortBy.Set(input)).Should(BeNil())
			Ω(sortBy).Should(Equal(output))
		},
		Entry(
			entities.SortByValues[0]+" case",
			entities.RecentlyOpenedSort,
			entities.SortByValues[0],
		),
		Entry(
			entities.SortByValues[1]+" case",
			entities.RecentlyImportedSort,
			entities.SortByValues[1],
		),
		Entry(
			entities.SortByValues[2]+" case",
			entities.OldestSort,
			entities.SortByValues[2],
		),
		Entry(
			entities.SortByValues[3]+" case",
			entities.NewestSort,
			entities.SortByValues[3],
		),
		Entry(
			entities.SortByValues[4]+" case",
			entities.LessDifficultSort,
			entities.SortByValues[4],
		),
		Entry(
			entities.SortByValues[5]+" case",
			entities.AlphabeticalSort,
			entities.SortByValues[5],
		),
	)
})
