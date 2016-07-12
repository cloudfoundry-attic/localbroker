package localbroker_test

import (
	"code.cloudfoundry.org/lager/lagertest"
	"github.com/cloudfoundry-incubator/localbroker/localbroker"
	"github.com/pivotal-cf/brokerapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var broker brokerapi.ServiceBroker

var _ = Describe("Broker", func() {
	BeforeEach(func() {
		logger := lagertest.NewTestLogger("test-broker")
		broker = localbroker.New(
			logger,
			"service-name", "service-id",
			"plan-name", "plan-id", "plan-desc",
		)
	})

	Context(".Services", func() {
		It("returns the service catalog as appropriate", func() {
			result := broker.Services()[0]
			Expect(result.ID).To(Equal("service-id"))
			Expect(result.Name).To(Equal("service-name"))
			Expect(result.Description).To(Equal("Local service docs: https://github.com/cloudfoundry-incubator/local-volume-release/"))
			Expect(result.Bindable).To(Equal(true))
			Expect(result.PlanUpdatable).To(Equal(false))
			Expect(result.Tags).To(ContainElement("local"))
			Expect(result.Requires).To(ContainElement(brokerapi.RequiredPermission("volume_mount")))

			Expect(result.Plans[0].Name).To(Equal("plan-name"))
			Expect(result.Plans[0].ID).To(Equal("plan-id"))
			Expect(result.Plans[0].Description).To(Equal("plan-desc"))
		})
	})
})
