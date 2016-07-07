// This file was generated by counterfeiter
package cephfakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/cephbroker/cephbrokerlocal"
	"github.com/cloudfoundry-incubator/cephbroker/model"
	"code.cloudfoundry.org/lager"
)

type FakeController struct {
	GetCatalogStub        func(logger lager.Logger) (model.Catalog, error)
	getCatalogMutex       sync.RWMutex
	getCatalogArgsForCall []struct {
		logger lager.Logger
	}
	getCatalogReturns struct {
		result1 model.Catalog
		result2 error
	}
	CreateServiceInstanceStub        func(logger lager.Logger, serverInstanceId string, instance model.ServiceInstance) (model.CreateServiceInstanceResponse, error)
	createServiceInstanceMutex       sync.RWMutex
	createServiceInstanceArgsForCall []struct {
		logger           lager.Logger
		serverInstanceId string
		instance         model.ServiceInstance
	}
	createServiceInstanceReturns struct {
		result1 model.CreateServiceInstanceResponse
		result2 error
	}
	ServiceInstanceExistsStub        func(logger lager.Logger, serviceInstanceId string) bool
	serviceInstanceExistsMutex       sync.RWMutex
	serviceInstanceExistsArgsForCall []struct {
		logger            lager.Logger
		serviceInstanceId string
	}
	serviceInstanceExistsReturns struct {
		result1 bool
	}
	ServiceInstancePropertiesMatchStub        func(logger lager.Logger, serviceInstanceId string, instance model.ServiceInstance) bool
	serviceInstancePropertiesMatchMutex       sync.RWMutex
	serviceInstancePropertiesMatchArgsForCall []struct {
		logger            lager.Logger
		serviceInstanceId string
		instance          model.ServiceInstance
	}
	serviceInstancePropertiesMatchReturns struct {
		result1 bool
	}
	DeleteServiceInstanceStub        func(logger lager.Logger, serviceInstanceId string) error
	deleteServiceInstanceMutex       sync.RWMutex
	deleteServiceInstanceArgsForCall []struct {
		logger            lager.Logger
		serviceInstanceId string
	}
	deleteServiceInstanceReturns struct {
		result1 error
	}
	BindServiceInstanceStub        func(logger lager.Logger, serverInstanceId string, bindingId string, bindingInfo model.ServiceBinding) (model.CreateServiceBindingResponse, error)
	bindServiceInstanceMutex       sync.RWMutex
	bindServiceInstanceArgsForCall []struct {
		logger           lager.Logger
		serverInstanceId string
		bindingId        string
		bindingInfo      model.ServiceBinding
	}
	bindServiceInstanceReturns struct {
		result1 model.CreateServiceBindingResponse
		result2 error
	}
	ServiceBindingExistsStub        func(logger lager.Logger, serviceInstanceId string, bindingId string) bool
	serviceBindingExistsMutex       sync.RWMutex
	serviceBindingExistsArgsForCall []struct {
		logger            lager.Logger
		serviceInstanceId string
		bindingId         string
	}
	serviceBindingExistsReturns struct {
		result1 bool
	}
	ServiceBindingPropertiesMatchStub        func(logger lager.Logger, serviceInstanceId string, bindingId string, binding model.ServiceBinding) bool
	serviceBindingPropertiesMatchMutex       sync.RWMutex
	serviceBindingPropertiesMatchArgsForCall []struct {
		logger            lager.Logger
		serviceInstanceId string
		bindingId         string
		binding           model.ServiceBinding
	}
	serviceBindingPropertiesMatchReturns struct {
		result1 bool
	}
	GetBindingStub        func(logger lager.Logger, serviceInstanceId, bindingId string) (model.ServiceBinding, error)
	getBindingMutex       sync.RWMutex
	getBindingArgsForCall []struct {
		logger            lager.Logger
		serviceInstanceId string
		bindingId         string
	}
	getBindingReturns struct {
		result1 model.ServiceBinding
		result2 error
	}
	UnbindServiceInstanceStub        func(logger lager.Logger, serviceInstanceId string, bindingId string) error
	unbindServiceInstanceMutex       sync.RWMutex
	unbindServiceInstanceArgsForCall []struct {
		logger            lager.Logger
		serviceInstanceId string
		bindingId         string
	}
	unbindServiceInstanceReturns struct {
		result1 error
	}
	invocations map[string][][]interface{}
}

func (fake *FakeController) GetCatalog(logger lager.Logger) (model.Catalog, error) {
	fake.getCatalogMutex.Lock()
	fake.getCatalogArgsForCall = append(fake.getCatalogArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.guard("GetCatalog")
	fake.invocations["GetCatalog"] = append(fake.invocations["GetCatalog"], []interface{}{logger})
	fake.getCatalogMutex.Unlock()
	if fake.GetCatalogStub != nil {
		return fake.GetCatalogStub(logger)
	} else {
		return fake.getCatalogReturns.result1, fake.getCatalogReturns.result2
	}
}

func (fake *FakeController) GetCatalogCallCount() int {
	fake.getCatalogMutex.RLock()
	defer fake.getCatalogMutex.RUnlock()
	return len(fake.getCatalogArgsForCall)
}

func (fake *FakeController) GetCatalogArgsForCall(i int) lager.Logger {
	fake.getCatalogMutex.RLock()
	defer fake.getCatalogMutex.RUnlock()
	return fake.getCatalogArgsForCall[i].logger
}

func (fake *FakeController) GetCatalogReturns(result1 model.Catalog, result2 error) {
	fake.GetCatalogStub = nil
	fake.getCatalogReturns = struct {
		result1 model.Catalog
		result2 error
	}{result1, result2}
}

func (fake *FakeController) CreateServiceInstance(logger lager.Logger, serverInstanceId string, instance model.ServiceInstance) (model.CreateServiceInstanceResponse, error) {
	fake.createServiceInstanceMutex.Lock()
	fake.createServiceInstanceArgsForCall = append(fake.createServiceInstanceArgsForCall, struct {
		logger           lager.Logger
		serverInstanceId string
		instance         model.ServiceInstance
	}{logger, serverInstanceId, instance})
	fake.guard("CreateServiceInstance")
	fake.invocations["CreateServiceInstance"] = append(fake.invocations["CreateServiceInstance"], []interface{}{logger, serverInstanceId, instance})
	fake.createServiceInstanceMutex.Unlock()
	if fake.CreateServiceInstanceStub != nil {
		return fake.CreateServiceInstanceStub(logger, serverInstanceId, instance)
	} else {
		return fake.createServiceInstanceReturns.result1, fake.createServiceInstanceReturns.result2
	}
}

func (fake *FakeController) CreateServiceInstanceCallCount() int {
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	return len(fake.createServiceInstanceArgsForCall)
}

func (fake *FakeController) CreateServiceInstanceArgsForCall(i int) (lager.Logger, string, model.ServiceInstance) {
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	return fake.createServiceInstanceArgsForCall[i].logger, fake.createServiceInstanceArgsForCall[i].serverInstanceId, fake.createServiceInstanceArgsForCall[i].instance
}

func (fake *FakeController) CreateServiceInstanceReturns(result1 model.CreateServiceInstanceResponse, result2 error) {
	fake.CreateServiceInstanceStub = nil
	fake.createServiceInstanceReturns = struct {
		result1 model.CreateServiceInstanceResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeController) ServiceInstanceExists(logger lager.Logger, serviceInstanceId string) bool {
	fake.serviceInstanceExistsMutex.Lock()
	fake.serviceInstanceExistsArgsForCall = append(fake.serviceInstanceExistsArgsForCall, struct {
		logger            lager.Logger
		serviceInstanceId string
	}{logger, serviceInstanceId})
	fake.guard("ServiceInstanceExists")
	fake.invocations["ServiceInstanceExists"] = append(fake.invocations["ServiceInstanceExists"], []interface{}{logger, serviceInstanceId})
	fake.serviceInstanceExistsMutex.Unlock()
	if fake.ServiceInstanceExistsStub != nil {
		return fake.ServiceInstanceExistsStub(logger, serviceInstanceId)
	} else {
		return fake.serviceInstanceExistsReturns.result1
	}
}

func (fake *FakeController) ServiceInstanceExistsCallCount() int {
	fake.serviceInstanceExistsMutex.RLock()
	defer fake.serviceInstanceExistsMutex.RUnlock()
	return len(fake.serviceInstanceExistsArgsForCall)
}

func (fake *FakeController) ServiceInstanceExistsArgsForCall(i int) (lager.Logger, string) {
	fake.serviceInstanceExistsMutex.RLock()
	defer fake.serviceInstanceExistsMutex.RUnlock()
	return fake.serviceInstanceExistsArgsForCall[i].logger, fake.serviceInstanceExistsArgsForCall[i].serviceInstanceId
}

func (fake *FakeController) ServiceInstanceExistsReturns(result1 bool) {
	fake.ServiceInstanceExistsStub = nil
	fake.serviceInstanceExistsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeController) ServiceInstancePropertiesMatch(logger lager.Logger, serviceInstanceId string, instance model.ServiceInstance) bool {
	fake.serviceInstancePropertiesMatchMutex.Lock()
	fake.serviceInstancePropertiesMatchArgsForCall = append(fake.serviceInstancePropertiesMatchArgsForCall, struct {
		logger            lager.Logger
		serviceInstanceId string
		instance          model.ServiceInstance
	}{logger, serviceInstanceId, instance})
	fake.guard("ServiceInstancePropertiesMatch")
	fake.invocations["ServiceInstancePropertiesMatch"] = append(fake.invocations["ServiceInstancePropertiesMatch"], []interface{}{logger, serviceInstanceId, instance})
	fake.serviceInstancePropertiesMatchMutex.Unlock()
	if fake.ServiceInstancePropertiesMatchStub != nil {
		return fake.ServiceInstancePropertiesMatchStub(logger, serviceInstanceId, instance)
	} else {
		return fake.serviceInstancePropertiesMatchReturns.result1
	}
}

func (fake *FakeController) ServiceInstancePropertiesMatchCallCount() int {
	fake.serviceInstancePropertiesMatchMutex.RLock()
	defer fake.serviceInstancePropertiesMatchMutex.RUnlock()
	return len(fake.serviceInstancePropertiesMatchArgsForCall)
}

func (fake *FakeController) ServiceInstancePropertiesMatchArgsForCall(i int) (lager.Logger, string, model.ServiceInstance) {
	fake.serviceInstancePropertiesMatchMutex.RLock()
	defer fake.serviceInstancePropertiesMatchMutex.RUnlock()
	return fake.serviceInstancePropertiesMatchArgsForCall[i].logger, fake.serviceInstancePropertiesMatchArgsForCall[i].serviceInstanceId, fake.serviceInstancePropertiesMatchArgsForCall[i].instance
}

func (fake *FakeController) ServiceInstancePropertiesMatchReturns(result1 bool) {
	fake.ServiceInstancePropertiesMatchStub = nil
	fake.serviceInstancePropertiesMatchReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeController) DeleteServiceInstance(logger lager.Logger, serviceInstanceId string) error {
	fake.deleteServiceInstanceMutex.Lock()
	fake.deleteServiceInstanceArgsForCall = append(fake.deleteServiceInstanceArgsForCall, struct {
		logger            lager.Logger
		serviceInstanceId string
	}{logger, serviceInstanceId})
	fake.guard("DeleteServiceInstance")
	fake.invocations["DeleteServiceInstance"] = append(fake.invocations["DeleteServiceInstance"], []interface{}{logger, serviceInstanceId})
	fake.deleteServiceInstanceMutex.Unlock()
	if fake.DeleteServiceInstanceStub != nil {
		return fake.DeleteServiceInstanceStub(logger, serviceInstanceId)
	} else {
		return fake.deleteServiceInstanceReturns.result1
	}
}

func (fake *FakeController) DeleteServiceInstanceCallCount() int {
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	return len(fake.deleteServiceInstanceArgsForCall)
}

func (fake *FakeController) DeleteServiceInstanceArgsForCall(i int) (lager.Logger, string) {
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	return fake.deleteServiceInstanceArgsForCall[i].logger, fake.deleteServiceInstanceArgsForCall[i].serviceInstanceId
}

func (fake *FakeController) DeleteServiceInstanceReturns(result1 error) {
	fake.DeleteServiceInstanceStub = nil
	fake.deleteServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeController) BindServiceInstance(logger lager.Logger, serverInstanceId string, bindingId string, bindingInfo model.ServiceBinding) (model.CreateServiceBindingResponse, error) {
	fake.bindServiceInstanceMutex.Lock()
	fake.bindServiceInstanceArgsForCall = append(fake.bindServiceInstanceArgsForCall, struct {
		logger           lager.Logger
		serverInstanceId string
		bindingId        string
		bindingInfo      model.ServiceBinding
	}{logger, serverInstanceId, bindingId, bindingInfo})
	fake.guard("BindServiceInstance")
	fake.invocations["BindServiceInstance"] = append(fake.invocations["BindServiceInstance"], []interface{}{logger, serverInstanceId, bindingId, bindingInfo})
	fake.bindServiceInstanceMutex.Unlock()
	if fake.BindServiceInstanceStub != nil {
		return fake.BindServiceInstanceStub(logger, serverInstanceId, bindingId, bindingInfo)
	} else {
		return fake.bindServiceInstanceReturns.result1, fake.bindServiceInstanceReturns.result2
	}
}

func (fake *FakeController) BindServiceInstanceCallCount() int {
	fake.bindServiceInstanceMutex.RLock()
	defer fake.bindServiceInstanceMutex.RUnlock()
	return len(fake.bindServiceInstanceArgsForCall)
}

func (fake *FakeController) BindServiceInstanceArgsForCall(i int) (lager.Logger, string, string, model.ServiceBinding) {
	fake.bindServiceInstanceMutex.RLock()
	defer fake.bindServiceInstanceMutex.RUnlock()
	return fake.bindServiceInstanceArgsForCall[i].logger, fake.bindServiceInstanceArgsForCall[i].serverInstanceId, fake.bindServiceInstanceArgsForCall[i].bindingId, fake.bindServiceInstanceArgsForCall[i].bindingInfo
}

func (fake *FakeController) BindServiceInstanceReturns(result1 model.CreateServiceBindingResponse, result2 error) {
	fake.BindServiceInstanceStub = nil
	fake.bindServiceInstanceReturns = struct {
		result1 model.CreateServiceBindingResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeController) ServiceBindingExists(logger lager.Logger, serviceInstanceId string, bindingId string) bool {
	fake.serviceBindingExistsMutex.Lock()
	fake.serviceBindingExistsArgsForCall = append(fake.serviceBindingExistsArgsForCall, struct {
		logger            lager.Logger
		serviceInstanceId string
		bindingId         string
	}{logger, serviceInstanceId, bindingId})
	fake.guard("ServiceBindingExists")
	fake.invocations["ServiceBindingExists"] = append(fake.invocations["ServiceBindingExists"], []interface{}{logger, serviceInstanceId, bindingId})
	fake.serviceBindingExistsMutex.Unlock()
	if fake.ServiceBindingExistsStub != nil {
		return fake.ServiceBindingExistsStub(logger, serviceInstanceId, bindingId)
	} else {
		return fake.serviceBindingExistsReturns.result1
	}
}

func (fake *FakeController) ServiceBindingExistsCallCount() int {
	fake.serviceBindingExistsMutex.RLock()
	defer fake.serviceBindingExistsMutex.RUnlock()
	return len(fake.serviceBindingExistsArgsForCall)
}

func (fake *FakeController) ServiceBindingExistsArgsForCall(i int) (lager.Logger, string, string) {
	fake.serviceBindingExistsMutex.RLock()
	defer fake.serviceBindingExistsMutex.RUnlock()
	return fake.serviceBindingExistsArgsForCall[i].logger, fake.serviceBindingExistsArgsForCall[i].serviceInstanceId, fake.serviceBindingExistsArgsForCall[i].bindingId
}

func (fake *FakeController) ServiceBindingExistsReturns(result1 bool) {
	fake.ServiceBindingExistsStub = nil
	fake.serviceBindingExistsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeController) ServiceBindingPropertiesMatch(logger lager.Logger, serviceInstanceId string, bindingId string, binding model.ServiceBinding) bool {
	fake.serviceBindingPropertiesMatchMutex.Lock()
	fake.serviceBindingPropertiesMatchArgsForCall = append(fake.serviceBindingPropertiesMatchArgsForCall, struct {
		logger            lager.Logger
		serviceInstanceId string
		bindingId         string
		binding           model.ServiceBinding
	}{logger, serviceInstanceId, bindingId, binding})
	fake.guard("ServiceBindingPropertiesMatch")
	fake.invocations["ServiceBindingPropertiesMatch"] = append(fake.invocations["ServiceBindingPropertiesMatch"], []interface{}{logger, serviceInstanceId, bindingId, binding})
	fake.serviceBindingPropertiesMatchMutex.Unlock()
	if fake.ServiceBindingPropertiesMatchStub != nil {
		return fake.ServiceBindingPropertiesMatchStub(logger, serviceInstanceId, bindingId, binding)
	} else {
		return fake.serviceBindingPropertiesMatchReturns.result1
	}
}

func (fake *FakeController) ServiceBindingPropertiesMatchCallCount() int {
	fake.serviceBindingPropertiesMatchMutex.RLock()
	defer fake.serviceBindingPropertiesMatchMutex.RUnlock()
	return len(fake.serviceBindingPropertiesMatchArgsForCall)
}

func (fake *FakeController) ServiceBindingPropertiesMatchArgsForCall(i int) (lager.Logger, string, string, model.ServiceBinding) {
	fake.serviceBindingPropertiesMatchMutex.RLock()
	defer fake.serviceBindingPropertiesMatchMutex.RUnlock()
	return fake.serviceBindingPropertiesMatchArgsForCall[i].logger, fake.serviceBindingPropertiesMatchArgsForCall[i].serviceInstanceId, fake.serviceBindingPropertiesMatchArgsForCall[i].bindingId, fake.serviceBindingPropertiesMatchArgsForCall[i].binding
}

func (fake *FakeController) ServiceBindingPropertiesMatchReturns(result1 bool) {
	fake.ServiceBindingPropertiesMatchStub = nil
	fake.serviceBindingPropertiesMatchReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeController) GetBinding(logger lager.Logger, serviceInstanceId string, bindingId string) (model.ServiceBinding, error) {
	fake.getBindingMutex.Lock()
	fake.getBindingArgsForCall = append(fake.getBindingArgsForCall, struct {
		logger            lager.Logger
		serviceInstanceId string
		bindingId         string
	}{logger, serviceInstanceId, bindingId})
	fake.guard("GetBinding")
	fake.invocations["GetBinding"] = append(fake.invocations["GetBinding"], []interface{}{logger, serviceInstanceId, bindingId})
	fake.getBindingMutex.Unlock()
	if fake.GetBindingStub != nil {
		return fake.GetBindingStub(logger, serviceInstanceId, bindingId)
	} else {
		return fake.getBindingReturns.result1, fake.getBindingReturns.result2
	}
}

func (fake *FakeController) GetBindingCallCount() int {
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	return len(fake.getBindingArgsForCall)
}

func (fake *FakeController) GetBindingArgsForCall(i int) (lager.Logger, string, string) {
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	return fake.getBindingArgsForCall[i].logger, fake.getBindingArgsForCall[i].serviceInstanceId, fake.getBindingArgsForCall[i].bindingId
}

func (fake *FakeController) GetBindingReturns(result1 model.ServiceBinding, result2 error) {
	fake.GetBindingStub = nil
	fake.getBindingReturns = struct {
		result1 model.ServiceBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeController) UnbindServiceInstance(logger lager.Logger, serviceInstanceId string, bindingId string) error {
	fake.unbindServiceInstanceMutex.Lock()
	fake.unbindServiceInstanceArgsForCall = append(fake.unbindServiceInstanceArgsForCall, struct {
		logger            lager.Logger
		serviceInstanceId string
		bindingId         string
	}{logger, serviceInstanceId, bindingId})
	fake.guard("UnbindServiceInstance")
	fake.invocations["UnbindServiceInstance"] = append(fake.invocations["UnbindServiceInstance"], []interface{}{logger, serviceInstanceId, bindingId})
	fake.unbindServiceInstanceMutex.Unlock()
	if fake.UnbindServiceInstanceStub != nil {
		return fake.UnbindServiceInstanceStub(logger, serviceInstanceId, bindingId)
	} else {
		return fake.unbindServiceInstanceReturns.result1
	}
}

func (fake *FakeController) UnbindServiceInstanceCallCount() int {
	fake.unbindServiceInstanceMutex.RLock()
	defer fake.unbindServiceInstanceMutex.RUnlock()
	return len(fake.unbindServiceInstanceArgsForCall)
}

func (fake *FakeController) UnbindServiceInstanceArgsForCall(i int) (lager.Logger, string, string) {
	fake.unbindServiceInstanceMutex.RLock()
	defer fake.unbindServiceInstanceMutex.RUnlock()
	return fake.unbindServiceInstanceArgsForCall[i].logger, fake.unbindServiceInstanceArgsForCall[i].serviceInstanceId, fake.unbindServiceInstanceArgsForCall[i].bindingId
}

func (fake *FakeController) UnbindServiceInstanceReturns(result1 error) {
	fake.UnbindServiceInstanceStub = nil
	fake.unbindServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeController) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeController) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ cephbrokerlocal.Controller = new(FakeController)
