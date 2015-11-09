// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
)

type FakeOrganizationRequirement struct {
	ExecuteStub        func() (success bool)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct{}
	executeReturns struct {
		result1 bool
	}
	SetOrganizationNameStub        func(string)
	setOrganizationNameMutex       sync.RWMutex
	setOrganizationNameArgsForCall []struct {
		arg1 string
	}
	GetOrganizationStub        func() models.Organization
	getOrganizationMutex       sync.RWMutex
	getOrganizationArgsForCall []struct{}
	getOrganizationReturns struct {
		result1 models.Organization
	}
}

func (fake *FakeOrganizationRequirement) Execute() (success bool) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeOrganizationRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeOrganizationRequirement) ExecuteReturns(result1 bool) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOrganizationRequirement) SetOrganizationName(arg1 string) {
	fake.setOrganizationNameMutex.Lock()
	fake.setOrganizationNameArgsForCall = append(fake.setOrganizationNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.setOrganizationNameMutex.Unlock()
	if fake.SetOrganizationNameStub != nil {
		fake.SetOrganizationNameStub(arg1)
	}
}

func (fake *FakeOrganizationRequirement) SetOrganizationNameCallCount() int {
	fake.setOrganizationNameMutex.RLock()
	defer fake.setOrganizationNameMutex.RUnlock()
	return len(fake.setOrganizationNameArgsForCall)
}

func (fake *FakeOrganizationRequirement) SetOrganizationNameArgsForCall(i int) string {
	fake.setOrganizationNameMutex.RLock()
	defer fake.setOrganizationNameMutex.RUnlock()
	return fake.setOrganizationNameArgsForCall[i].arg1
}

func (fake *FakeOrganizationRequirement) GetOrganization() models.Organization {
	fake.getOrganizationMutex.Lock()
	fake.getOrganizationArgsForCall = append(fake.getOrganizationArgsForCall, struct{}{})
	fake.getOrganizationMutex.Unlock()
	if fake.GetOrganizationStub != nil {
		return fake.GetOrganizationStub()
	} else {
		return fake.getOrganizationReturns.result1
	}
}

func (fake *FakeOrganizationRequirement) GetOrganizationCallCount() int {
	fake.getOrganizationMutex.RLock()
	defer fake.getOrganizationMutex.RUnlock()
	return len(fake.getOrganizationArgsForCall)
}

func (fake *FakeOrganizationRequirement) GetOrganizationReturns(result1 models.Organization) {
	fake.GetOrganizationStub = nil
	fake.getOrganizationReturns = struct {
		result1 models.Organization
	}{result1}
}

var _ requirements.OrganizationRequirement = new(FakeOrganizationRequirement)
