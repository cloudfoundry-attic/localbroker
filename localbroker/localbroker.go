package localbroker

import (
	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
)

const PermissionVolumeMount = brokerapi.RequiredPermission("volume_mount")

type broker struct {
	logger      lager.Logger
	serviceName string
	serviceId   string

	planName string
	planId   string
	planDesc string

	// instanceMap map[string]*model.ServiceInstance
	// bindingMap  map[string]*model.ServiceBinding
}

func New(
	logger lager.Logger,
	serviceName, serviceId, planName, planId, planDesc string,
	// instanceMap map[string]*model.ServiceInstance,
	// bindingMap map[string]*model.ServiceBinding,
) *broker {
	return &broker{
		logger:      logger,
		serviceName: serviceName,
		serviceId:   serviceId,
		planName:    planName,
		planId:      planId,
		planDesc:    planDesc,
		// instanceMap: instanceMap,
		// bindingMap:  bindingMap,
	}
}

func (b *broker) Services() []brokerapi.Service {
	logger := b.logger.Session("get-catalog")
	logger.Info("start")
	defer logger.Info("end")

	return []brokerapi.Service{{
		ID:            b.serviceId,
		Name:          b.serviceName,
		Description:   "Local service docs: https://github.com/cloudfoundry-incubator/local-volume-release/",
		Bindable:      true,
		PlanUpdatable: false,
		Tags:          []string{"local"},
		Requires:      []brokerapi.RequiredPermission{PermissionVolumeMount},

		Plans: []brokerapi.ServicePlan{{
			Name:        b.planName,
			ID:          b.planId,
			Description: b.planDesc,
		}},
	}}
}

func (b *broker) Provision(instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	panic("not implemented")
}

func (b *broker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	panic("not implemented")
}

func (b *broker) Bind(instanceID string, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	panic("not implemented")
}

func (b *broker) Unbind(instanceID string, bindingID string, details brokerapi.UnbindDetails) error {
	panic("not implemented")
}

func (b *broker) Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	panic("not implemented")
}

func (b *broker) LastOperation(instanceID string, operationData string) (brokerapi.LastOperation, error) {
	panic("not implemented")
}
