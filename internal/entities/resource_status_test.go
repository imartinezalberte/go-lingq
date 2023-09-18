package entities_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go-lingq/internal/entities"
)

var _ = Describe("ResourceStatus", func() {
	var resourceStatus entities.ResourceStatus

	When("Input is not correct", func() {
		It("should return error", func() {
			Ω(resourceStatus.Set("incorrect input")).Should(MatchError("unknown resource status"))
		})
	})

	When("Content is incorrect", func() {
		BeforeEach(func() {
			resourceStatus = entities.ResourceStatus(len(entities.ResourceStatusValue))
		})

		It("should return \"unknown\" string", func() {
			Ω(resourceStatus.String()).Should(Equal("unknown"))
		})
	})

	DescribeTable(
		"String should return the right type",
		func(input entities.ResourceStatus, output string) {
			Ω(input.String()).Should(Equal(output))
		},
		Entry("private", entities.PrivateResource, entities.ResourceStatusValue[0]),
		Entry("shared", entities.SharedResource, entities.ResourceStatusValue[1]),
	)

	DescribeTable(
		"Set works with every possible option",
		func(output entities.ResourceStatus, input string) {
			Ω(resourceStatus.Set(input)).Should(BeNil())
			Ω(resourceStatus).Should(Equal(output))
		},
		Entry(
			entities.ResourceStatusValue[0]+" case",
			entities.PrivateResource,
			entities.ResourceStatusValue[0],
		),
		Entry(
			entities.ResourceStatusValue[1]+" case",
			entities.SharedResource,
			entities.ResourceStatusValue[1],
		),
	)
})
