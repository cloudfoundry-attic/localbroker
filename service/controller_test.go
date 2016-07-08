package service_test

import (
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	"github.com/cloudfoundry-incubator/localbroker/model"
	. "github.com/cloudfoundry-incubator/localbroker/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("localbroker service", func() {
	var (
		testLogger  lager.Logger
		controller  Controller
		instanceMap map[string]*model.ServiceInstance
		bindingMap  map[string]*model.ServiceBinding
		planId      string
		planName    string
		planDesc    string
	)

	BeforeEach(func() {
		planName = "free"
		planId = "free-plan-guid"
		planDesc = "free local filesystem"
		testLogger = lagertest.NewTestLogger("ControllerTest")
		instanceMap = make(map[string]*model.ServiceInstance)
		bindingMap = make(map[string]*model.ServiceBinding)
		controller = NewController("service-name", "service-id", planId, planName, planDesc, "/tmp/localbroker", instanceMap, bindingMap)
	})

	Context(".Catalog", func() {
		It("should produce a valid catalog", func() {
			catalog, err := controller.GetCatalog(testLogger)
			Expect(err).ToNot(HaveOccurred())

			Expect(catalog).ToNot(BeNil())
			Expect(catalog.Services).ToNot(BeNil())
			Expect(len(catalog.Services)).To(Equal(1))

			Expect(catalog.Services[0].Name).To(Equal("service-name"))
			Expect(catalog.Services[0].Id).To(Equal("service-id"))

			Expect(catalog.Services[0].Requires).ToNot(BeNil())
			Expect(len(catalog.Services[0].Requires)).To(Equal(1))
			Expect(catalog.Services[0].Requires[0]).To(Equal("volume_mount"))

			Expect(catalog.Services[0].Plans).ToNot(BeNil())
			Expect(len(catalog.Services[0].Plans)).To(Equal(1))
			Expect(catalog.Services[0].Plans[0].Name).To(Equal(planName))

			Expect(catalog.Services[0].Bindable).To(Equal(true))
			Expect(catalog.Services[0].PlanUpdateable).To(Equal(false))
		})
	})
})
