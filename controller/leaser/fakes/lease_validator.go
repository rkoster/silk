// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/silk/controller"
)

type LeaseValidator struct {
	ValidateStub        func(controller.Lease) error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		arg1 controller.Lease
	}
	validateReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LeaseValidator) Validate(arg1 controller.Lease) error {
	fake.validateMutex.Lock()
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		arg1 controller.Lease
	}{arg1})
	fake.recordInvocation("Validate", []interface{}{arg1})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub(arg1)
	}
	return fake.validateReturns.result1
}

func (fake *LeaseValidator) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *LeaseValidator) ValidateArgsForCall(i int) controller.Lease {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.validateArgsForCall[i].arg1
}

func (fake *LeaseValidator) ValidateReturns(result1 error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *LeaseValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.invocations
}

func (fake *LeaseValidator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
