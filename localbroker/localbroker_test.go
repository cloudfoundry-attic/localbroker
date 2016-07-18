package localbroker_test

import (
	"code.cloudfoundry.org/lager/lagertest"
	"github.com/cloudfoundry-incubator/localbroker/localbroker"
	"github.com/cloudfoundry-incubator/voldriver"
	"github.com/cloudfoundry-incubator/voldriver/voldriverfakes"
	"github.com/pivotal-cf/brokerapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Broker", func() {
	var (
		broker          brokerapi.ServiceBroker
		fakeProvisioner *voldriverfakes.FakeProvisioner
	)

	BeforeEach(func() {
		logger := lagertest.NewTestLogger("test-broker")
		fakeProvisioner = &voldriverfakes.FakeProvisioner{}
		broker = localbroker.New(
			logger, fakeProvisioner,
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

	Context(".Provision", func() {
		It("should provision the service instance", func() {
			_, err := broker.Provision("some-instance-id", brokerapi.ProvisionDetails{}, false)
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeProvisioner.CreateCallCount()).To(Equal(1))

			_, details := fakeProvisioner.CreateArgsForCall(0)
			Expect(err).NotTo(HaveOccurred())
			Expect(details.Name).To(Equal("some-instance-id"))
			Expect(details.Opts["volume_id"]).To(Equal("some-instance-id"))
		})

		Context("when provisioning errors", func() {
			BeforeEach(func() {
				fakeProvisioner.CreateReturns(voldriver.ErrorResponse{Err: "some-error"})
			})

			It("errors", func() {
				_, err := broker.Provision("some-instance-id", brokerapi.ProvisionDetails{}, false)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when the service instance already exists with different details", func() {
			var details brokerapi.ProvisionDetails
			BeforeEach(func() {
				details = brokerapi.ProvisionDetails{
					ServiceID:        "service-id",
					PlanID:           "plan-id",
					OrganizationGUID: "org-guid",
					SpaceGUID:        "space-guid",
				}
				_, err := broker.Provision("some-instance-id", details, false)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should error", func() {
				details.ServiceID = "different-service-id"
				_, err := broker.Provision("some-instance-id", details, false)
				Expect(err).To(Equal(brokerapi.ErrInstanceAlreadyExists))
			})
		})
	})

	Context(".Deprovision", func() {
		BeforeEach(func() {
			_, err := broker.Provision("some-instance-id", brokerapi.ProvisionDetails{}, false)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should deprovision the service", func() {
			_, err := broker.Deprovision("some-instance-id", brokerapi.DeprovisionDetails{}, false)
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeProvisioner.RemoveCallCount()).To(Equal(1))

			By("checking that we can reprovision a slightly different service")
			_, err = broker.Provision("some-instance-id", brokerapi.ProvisionDetails{ServiceID: "different-service-id"}, false)
			Expect(err).NotTo(Equal(brokerapi.ErrInstanceAlreadyExists))
		})

		It("errors when the service instance does not exist", func() {
			_, err := broker.Deprovision("some-nonexistant-instance-id", brokerapi.DeprovisionDetails{}, false)
			Expect(err).To(Equal(brokerapi.ErrInstanceDoesNotExist))
		})

		Context("when the provisioner fails to remove", func() {
			BeforeEach(func() {
				fakeProvisioner.RemoveReturns(voldriver.ErrorResponse{Err: "some-error"})
			})

			It("should error", func() {
				_, err := broker.Deprovision("some-instance-id", brokerapi.DeprovisionDetails{}, false)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context(".Bind", func() {
		var bindDetails brokerapi.BindDetails

		BeforeEach(func() {
			_, err := broker.Provision("some-instance-id", brokerapi.ProvisionDetails{}, false)
			Expect(err).NotTo(HaveOccurred())

			bindDetails = brokerapi.BindDetails{AppGUID: "guid", Parameters: map[string]interface{}{}}
		})

		It("includes empty credentials to prevent CAPI crash", func() {
			binding, err := broker.Bind("some-instance-id", "binding-id", bindDetails)
			Expect(err).NotTo(HaveOccurred())

			Expect(binding.Credentials).NotTo(BeNil())
		})

		It("uses the instance id in the default container path", func() {
			binding, err := broker.Bind("some-instance-id", "binding-id", bindDetails)
			Expect(err).NotTo(HaveOccurred())
			Expect(binding.VolumeMounts[0].ContainerPath).To(Equal("/var/vcap/data/some-instance-id"))
		})

		It("flows container path through", func() {
			bindDetails.Parameters["mount"] = "/var/vcap/otherdir/something"
			binding, err := broker.Bind("some-instance-id", "binding-id", bindDetails)
			Expect(err).NotTo(HaveOccurred())
			Expect(binding.VolumeMounts[0].ContainerPath).To(Equal("/var/vcap/otherdir/something"))
		})

		It("uses rw as its default mode", func() {
			binding, err := broker.Bind("some-instance-id", "binding-id", bindDetails)
			Expect(err).NotTo(HaveOccurred())
			Expect(binding.VolumeMounts[0].Mode).To(Equal("rw"))
		})

		It("sets mode to `r` when readonly is true", func() {
			bindDetails.Parameters["readonly"] = true
			binding, err := broker.Bind("some-instance-id", "binding-id", bindDetails)
			Expect(err).NotTo(HaveOccurred())

			Expect(binding.VolumeMounts[0].Mode).To(Equal("r"))
		})

		It("errors if mode is not a boolean", func() {
			bindDetails.Parameters["readonly"] = ""
			_, err := broker.Bind("some-instance-id", "binding-id", bindDetails)
			Expect(err).To(Equal(brokerapi.ErrRawParamsInvalid))
		})

		It("fills in the driver name", func() {
			binding, err := broker.Bind("some-instance-id", "binding-id", bindDetails)
			Expect(err).NotTo(HaveOccurred())

			Expect(binding.VolumeMounts[0].Private.Driver).To(Equal("localdriver"))
		})

		It("fills in the group id", func() {
			binding, err := broker.Bind("some-instance-id", "binding-id", bindDetails)
			Expect(err).NotTo(HaveOccurred())

			Expect(binding.VolumeMounts[0].Private.GroupId).To(Equal("some-instance-id"))
		})

		Context("when the binding already exists", func() {
			BeforeEach(func() {
				_, err := broker.Bind("some-instance-id", "binding-id", brokerapi.BindDetails{AppGUID: "guid"})
				Expect(err).NotTo(HaveOccurred())
			})

			It("doesn't error when binding the same details", func() {
				_, err := broker.Bind("some-instance-id", "binding-id", brokerapi.BindDetails{AppGUID: "guid"})
				Expect(err).NotTo(HaveOccurred())
			})

			It("errors when binding different details", func() {
				_, err := broker.Bind("some-instance-id", "binding-id", brokerapi.BindDetails{AppGUID: "different"})
				Expect(err).To(Equal(brokerapi.ErrBindingAlreadyExists))
			})
		})

		It("errors when the service instance does not exist", func() {
			_, err := broker.Bind("nonexistant-instance-id", "binding-id", brokerapi.BindDetails{AppGUID: "guid"})
			Expect(err).To(Equal(brokerapi.ErrInstanceDoesNotExist))
		})

		It("errors when the app guid is not provided", func() {
			_, err := broker.Bind("some-instance-id", "binding-id", brokerapi.BindDetails{})
			Expect(err).To(Equal(brokerapi.ErrAppGuidNotProvided))
		})
	})

	Context(".Unbind", func() {
		BeforeEach(func() {
			_, err := broker.Provision("some-instance-id", brokerapi.ProvisionDetails{}, false)
			Expect(err).NotTo(HaveOccurred())

			_, err = broker.Bind("some-instance-id", "binding-id", brokerapi.BindDetails{AppGUID: "guid"})
			Expect(err).NotTo(HaveOccurred())
		})

		It("unbinds a bound service instance from an app", func() {
			err := broker.Unbind("some-instance-id", "binding-id", brokerapi.UnbindDetails{})
			Expect(err).NotTo(HaveOccurred())
		})

		It("fails when trying to unbind a instance that has not been provisioned", func() {
			err := broker.Unbind("some-other-instance-id", "binding-id", brokerapi.UnbindDetails{})
			Expect(err).To(Equal(brokerapi.ErrInstanceDoesNotExist))
		})

		It("fails when trying to unbind a binding that has not been bound", func() {
			err := broker.Unbind("some-instance-id", "some-other-binding-id", brokerapi.UnbindDetails{})
			Expect(err).To(Equal(brokerapi.ErrBindingDoesNotExist))
		})
	})
})