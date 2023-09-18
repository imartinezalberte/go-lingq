package entities_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go-lingq/internal/entities"
)

var _ = Describe("ResourceLevel", func() {
	var resourceLevel entities.ResourceLevel

	When("Input is not correct", func() {
		It("should return error", func() {
			Ω(resourceLevel.Set("incorrect input")).Should(MatchError("unknown resource level"))
		})
	})

	When("Content is incorrect", func() {
		BeforeEach(func() {
			resourceLevel = entities.ResourceLevel(len(entities.Levels))
		})

		It("should return \"unknown\" string", func() {
			Ω(resourceLevel.String()).Should(Equal("unknown"))
		})
	})

	DescribeTable(
		"String should return the right type",
		func(input entities.ResourceLevel, output string) {
			Ω(input.String()).Should(Equal(output))
		},
		Entry("first level", entities.FirstLevel, strings.ToUpper(entities.Levels[0])),
		Entry("second level", entities.SecondLevel, strings.ToUpper(entities.Levels[1])),
		Entry("third level", entities.ThirdLevel, strings.ToUpper(entities.Levels[2])),
		Entry("fourth level", entities.FourthLevel, strings.ToUpper(entities.Levels[3])),
		Entry("fifth level", entities.FifthLevel, strings.ToUpper(entities.Levels[4])),
		Entry("sixth level", entities.SixthLevel, strings.ToUpper(entities.Levels[5])),
	)

	DescribeTable(
		"Set works with every possible option",
		func(output entities.ResourceLevel, possibleInputs ...string) {
			Ω(resourceLevel.Set(possibleInputs[0])).Should(BeNil())
			Ω(resourceLevel).Should(Equal(output))

			Ω(resourceLevel.Set(possibleInputs[1])).Should(BeNil())
			Ω(resourceLevel).Should(Equal(output))
		},
		Entry(
			entities.UnmarshalLevels[0]+" case",
			entities.FirstLevel,
			entities.Levels[0],
			entities.UnmarshalLevels[0],
		),
		Entry(
			entities.UnmarshalLevels[1]+" case",
			entities.SecondLevel,
			entities.Levels[1],
			entities.UnmarshalLevels[1],
		),
		Entry(
			entities.UnmarshalLevels[2]+" case",
			entities.ThirdLevel,
			entities.Levels[2],
			entities.UnmarshalLevels[2],
		),
		Entry(
			entities.UnmarshalLevels[3]+" case",
			entities.FourthLevel,
			entities.Levels[3],
			entities.UnmarshalLevels[3],
		),
		Entry(
			entities.UnmarshalLevels[4]+" case",
			entities.FifthLevel,
			entities.Levels[4],
			entities.UnmarshalLevels[4],
		),
		Entry(
			entities.UnmarshalLevels[5]+" case",
			entities.SixthLevel,
			entities.Levels[5],
			entities.UnmarshalLevels[5],
		),
	)
})
