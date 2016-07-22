package localbroker

import (
	"errors"
	"reflect"

	"path"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-incubator/voldriver"
	"github.com/pivotal-cf/brokerapi"
)

const (
	PermissionVolumeMount = brokerapi.RequiredPermission("volume_mount")
	DefaultContainerPath  = "/var/vcap/data"
)

type staticState struct {
	ServiceName string `json:"ServiceName"`
	ServiceId   string `json:"ServiceId"`
	PlanName    string `json:"PlanName"`
	PlanId      string `json:"PlanId"`
	PlanDesc    string `json:"PlanDesc"`
}

type dynamicState struct {
	InstanceMap map[string]brokerapi.ProvisionDetails
	BindingMap  map[string]brokerapi.BindDetails
}

type broker struct {
	logger      lager.Logger
	provisioner voldriver.Provisioner
	dataDir     string
	fs          FileSystem
	static      staticState
	dynamic     dynamicState
}

func New(
	logger lager.Logger, provisioner voldriver.Provisioner,
	serviceName, serviceId, planName, planId, planDesc, dataDir string,
	fileSystem FileSystem,
) *broker {

	return &broker{
		logger:      logger,
		provisioner: provisioner,
		dataDir:     dataDir,
		fs:          fileSystem,
		static: staticState{
			ServiceName: serviceName,
			ServiceId:   serviceId,
			PlanName:    planName,
			PlanId:      planId,
			PlanDesc:    planDesc,
		},
		dynamic: dynamicState{
			InstanceMap: map[string]brokerapi.ProvisionDetails{},
			BindingMap:  map[string]brokerapi.BindDetails{},
		},
	}
}

func (b *broker) Services() []brokerapi.Service {
	logger := b.logger.Session("services")
	logger.Info("start")
	defer logger.Info("end")

	return []brokerapi.Service{{
		ID:            b.static.ServiceId,
		Name:          b.static.ServiceName,
		Description:   "Local service docs: https://github.com/cloudfoundry-incubator/local-volume-release/",
		Bindable:      true,
		PlanUpdatable: false,
		Tags:          []string{"local"},
		Requires:      []brokerapi.RequiredPermission{PermissionVolumeMount},

		Plans: []brokerapi.ServicePlan{{
			Name:        b.static.PlanName,
			ID:          b.static.PlanId,
			Description: b.static.PlanDesc,
		}},
	}}
}

func (b *broker) Provision(instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	logger := b.logger.Session("provision")
	logger.Info("start")
	defer logger.Info("end")

	if b.instanceConflicts(details, instanceID) {
		logger.Error("instance-already-exists", brokerapi.ErrInstanceAlreadyExists)
		return brokerapi.ProvisionedServiceSpec{}, brokerapi.ErrInstanceAlreadyExists
	}

	errResp := b.provisioner.Create(logger, voldriver.CreateRequest{
		Name: instanceID,
		Opts: map[string]interface{}{"volume_id": instanceID},
	})

	if errResp.Err != "" {
		err := errors.New(errResp.Err)
		logger.Error("provisioner-create-failed", err)
		return brokerapi.ProvisionedServiceSpec{}, err
	}

	b.dynamic.InstanceMap[instanceID] = details

	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (b *broker) instanceConflicts(details brokerapi.ProvisionDetails, instanceID string) bool {
	if existing, ok := b.dynamic.InstanceMap[instanceID]; ok {
		if !reflect.DeepEqual(details, existing) {
			return true
		}
	}
	return false
}

func (b *broker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	logger := b.logger.Session("deprovision")
	logger.Info("start")
	defer logger.Info("end")

	if _, ok := b.dynamic.InstanceMap[instanceID]; !ok {
		return brokerapi.DeprovisionServiceSpec{}, brokerapi.ErrInstanceDoesNotExist
	}

	errResp := b.provisioner.Remove(logger, voldriver.RemoveRequest{
		Name: instanceID,
	})

	if errResp.Err != "" {
		err := errors.New(errResp.Err)
		logger.Error("provisioner-remove-failed", err)
		return brokerapi.DeprovisionServiceSpec{}, err
	}

	delete(b.dynamic.InstanceMap, instanceID)

	return brokerapi.DeprovisionServiceSpec{}, nil
}

func (b *broker) Bind(instanceID string, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	if _, ok := b.dynamic.InstanceMap[instanceID]; !ok {
		return brokerapi.Binding{}, brokerapi.ErrInstanceDoesNotExist
	}

	if details.AppGUID == "" {
		return brokerapi.Binding{}, brokerapi.ErrAppGuidNotProvided
	}

	mode, err := evaluateMode(details.Parameters)
	if err != nil {
		return brokerapi.Binding{}, err
	}

	if b.bindingConflicts(bindingID, details) {
		return brokerapi.Binding{}, brokerapi.ErrBindingAlreadyExists
	}

	b.dynamic.BindingMap[bindingID] = details

	return brokerapi.Binding{
		Credentials: struct{}{}, // if nil, cloud controller chokes on response
		VolumeMounts: []brokerapi.VolumeMount{{
			ContainerPath: evaluateContainerPath(details.Parameters, instanceID),
			Mode:          mode,
			Private: brokerapi.VolumeMountPrivate{
				Driver:  "localdriver",
				GroupId: instanceID,
			},
		}},
	}, nil
}

func evaluateContainerPath(parameters map[string]interface{}, volId string) string {
	if containerPath, ok := parameters["mount"]; ok && containerPath != "" {
		return containerPath.(string)
	}

	return path.Join(DefaultContainerPath, volId)
}

func evaluateMode(parameters map[string]interface{}) (string, error) {
	if ro, ok := parameters["readonly"]; ok {
		switch ro := ro.(type) {
		case bool:
			return readOnlyToMode(ro), nil
		default:
			return "", brokerapi.ErrRawParamsInvalid
		}
	}
	return "rw", nil
}

func readOnlyToMode(ro bool) string {
	if ro {
		return "r"
	}
	return "rw"
}

func (b *broker) Unbind(instanceID string, bindingID string, details brokerapi.UnbindDetails) error {
	if _, ok := b.dynamic.InstanceMap[instanceID]; !ok {
		return brokerapi.ErrInstanceDoesNotExist
	}

	if _, ok := b.dynamic.BindingMap[bindingID]; !ok {
		return brokerapi.ErrBindingDoesNotExist
	}

	delete(b.dynamic.BindingMap, bindingID)

	return nil
}

func (b *broker) bindingConflicts(bindingID string, details brokerapi.BindDetails) bool {
	if existing, ok := b.dynamic.BindingMap[bindingID]; ok {
		if !reflect.DeepEqual(details, existing) {
			return true
		}
	}
	return false
}

func (b *broker) Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	panic("not implemented")
}

func (b *broker) LastOperation(instanceID string, operationData string) (brokerapi.LastOperation, error) {
	panic("not implemented")
}
