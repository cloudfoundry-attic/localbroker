package service

import (
	"errors"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-incubator/localbroker/model"
)

const (
	DEFAULT_POLLING_INTERVAL_SECONDS = 10
	DEFAULT_CONTAINER_PATH           = "/var/vcap/data/"
)

//go:generate counterfeiter -o ./servicefakes/fake_controller.go . Controller

type Controller interface {
	GetCatalog(logger lager.Logger) (model.Catalog, error)
	CreateServiceInstance(logger lager.Logger, serverInstanceId string, instance model.ServiceInstance) (model.CreateServiceInstanceResponse, error)
	ServiceInstanceExists(logger lager.Logger, serviceInstanceId string) bool
	ServiceInstancePropertiesMatch(logger lager.Logger, serviceInstanceId string, instance model.ServiceInstance) bool
	DeleteServiceInstance(logger lager.Logger, serviceInstanceId string) error
	BindServiceInstance(logger lager.Logger, serverInstanceId string, bindingId string, bindingInfo model.ServiceBinding) (model.CreateServiceBindingResponse, error)
	ServiceBindingExists(logger lager.Logger, serviceInstanceId string, bindingId string) bool
	ServiceBindingPropertiesMatch(logger lager.Logger, serviceInstanceId string, bindingId string, binding model.ServiceBinding) bool
	GetBinding(logger lager.Logger, serviceInstanceId, bindingId string) (model.ServiceBinding, error)
	UnbindServiceInstance(logger lager.Logger, serviceInstanceId string, bindingId string) error
}

type brokerController struct {
	instanceMap map[string]*model.ServiceInstance
	bindingMap  map[string]*model.ServiceBinding
	configPath  string
	serviceName string
	serviceId   string
	planId      string
	planName    string
	planDesc    string
}

func NewController(serviceName, serviceId, planId, planName, planDesc, configPath string, instanceMap map[string]*model.ServiceInstance, bindingMap map[string]*model.ServiceBinding) Controller {
	return &brokerController{
		serviceName: serviceName,
		serviceId:   serviceId,
		planId:      planId,
		planName:    planName,
		planDesc:    planDesc,
		configPath:  configPath,
		instanceMap: instanceMap,
		bindingMap:  bindingMap,
	}
}

func (c *brokerController) GetCatalog(logger lager.Logger) (model.Catalog, error) {
	logger = logger.Session("get-catalog")
	logger.Info("start")
	defer logger.Info("end")
	plan := model.ServicePlan{
		Name:        c.planName, // "free"
		Id:          c.planId,   // "free-plan-guid"
		Description: c.planDesc,
		Metadata:    nil,
		Free:        true,
	}

	service := model.Service{
		Name:            c.serviceName,
		Id:              c.serviceId,
		Description:     "Local service docs: https://github.com/cloudfoundry-incubator/local-volume-release/",
		Bindable:        true,
		PlanUpdateable:  false,
		Tags:            []string{"local"},
		Requires:        []string{"volume_mount"},
		Metadata:        nil,
		Plans:           []model.ServicePlan{plan},
		DashboardClient: nil,
	}
	catalog := model.Catalog{
		Services: []model.Service{service},
	}
	return catalog, nil
}

func (c *brokerController) CreateServiceInstance(logger lager.Logger, serviceInstanceId string, instance model.ServiceInstance) (model.CreateServiceInstanceResponse, error) {
	logger = logger.Session("create-service-instance")
	logger.Info("start")
	defer logger.Info("end")

	return model.CreateServiceInstanceResponse{}, errors.New("unimplemented")
}

func (c *brokerController) ServiceInstanceExists(logger lager.Logger, serviceInstanceId string) bool {
	logger = logger.Session("service-instance-exists")
	logger.Info("start")
	defer logger.Info("end")

	return false
}

func (c *brokerController) ServiceInstancePropertiesMatch(logger lager.Logger, serviceInstanceId string, instance model.ServiceInstance) bool {
	logger = logger.Session("service-instance-properties-match")
	logger.Info("start")
	defer logger.Info("end")

	return false
}

func (c *brokerController) DeleteServiceInstance(logger lager.Logger, serviceInstanceId string) error {
	logger = logger.Session("delete-service-instance")
	logger.Info("start")
	defer logger.Info("end")

	return errors.New("unimplemented")
}
func (c *brokerController) BindServiceInstance(logger lager.Logger, serviceInstanceId string, bindingId string, bindingInfo model.ServiceBinding) (model.CreateServiceBindingResponse, error) {
	logger = logger.Session("bind-service-instance")
	logger.Info("start")
	defer logger.Info("end")

	return model.CreateServiceBindingResponse{}, errors.New("unimplemented")
}

func (c *brokerController) ServiceBindingExists(logger lager.Logger, serviceInstanceId string, bindingId string) bool {
	logger = logger.Session("service-binding-exists")
	logger.Info("start")
	defer logger.Info("end")

	return false
}

func (c *brokerController) ServiceBindingPropertiesMatch(logger lager.Logger, serviceInstanceId string, bindingId string, binding model.ServiceBinding) bool {
	logger = logger.Session("service-binding-properties-match")
	logger.Info("start")
	defer logger.Info("end")

	return false
}

func (c *brokerController) UnbindServiceInstance(logger lager.Logger, serviceInstanceId string, bindingId string) error {
	logger = logger.Session("unbind")
	logger.Info("start")
	defer logger.Info("end")

	return errors.New("unimplemented")
}

func (c *brokerController) GetBinding(logger lager.Logger, instanceId, bindingId string) (model.ServiceBinding, error) {
	logger = logger.Session("get-binding")
	logger.Info("start")
	defer logger.Info("end")

	return model.ServiceBinding{}, errors.New("unimplemented")
}
