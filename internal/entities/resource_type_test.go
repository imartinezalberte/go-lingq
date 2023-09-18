package entities_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go-lingq/internal/entities"
)

var _ = Describe("ResourceType", func() {
	var resourceType entities.ResourceType

	When("Input is not correct", func() {
		It("should return error", func() {
			Ω(resourceType.Set("incorrect input")).Should(MatchError("unknown resource type"))
		})
	})

	When("Content is incorrect", func() {
		BeforeEach(func() {
			resourceType = entities.ResourceType(len(entities.ResourceTypeValues))
		})

		It("should return \"unknown\" string", func() {
			Ω(resourceType.String()).Should(Equal("unknown"))
		})
	})

	DescribeTable(
		"String should return the right type",
		func(input entities.ResourceType, output string) {
			Ω(input.String()).Should(Equal(output))
		},
		Entry("private", entities.Course, entities.ResourceTypeValues[0]),
		Entry("shared", entities.Lesson, entities.ResourceTypeValues[1]),
	)

	DescribeTable(
		"Set works with every possible option",
		func(output entities.ResourceType, possibleInputs ...string) {
			for _, input := range possibleInputs {
				Ω(resourceType.Set(input)).Should(BeNil())
				Ω(resourceType).Should(Equal(output))
			}
		},
		Entry(
			entities.ResourceTypeValues[0]+" case",
			entities.Course,
			entities.ResourceTypeValues[0],
			entities.ResourceTypeUnmarshalValues[0],
		),
		Entry(
			entities.ResourceTypeValues[1]+" case",
			entities.Lesson,
			entities.ResourceTypeValues[1],
			entities.ResourceTypeUnmarshalValues[1],
		),
	)
})
