// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/silk/controller"
)

type Converger struct {
	ConvergeStub        func([]controller.Lease) error
	convergeMutex       sync.RWMutex
	convergeArgsForCall []struct {
		arg1 []controller.Lease
	}
	convergeReturns struct {
		result1 error
	}
	convergeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Converger) Converge(arg1 []controller.Lease) error {
	var arg1Copy []controller.Lease
	if arg1 != nil {
		arg1Copy = make([]controller.Lease, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.convergeMutex.Lock()
	ret, specificReturn := fake.convergeReturnsOnCall[len(fake.convergeArgsForCall)]
	fake.convergeArgsForCall = append(fake.convergeArgsForCall, struct {
		arg1 []controller.Lease
	}{arg1Copy})
	fake.recordInvocation("Converge", []interface{}{arg1Copy})
	fake.convergeMutex.Unlock()
	if fake.ConvergeStub != nil {
		return fake.ConvergeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.convergeReturns.result1
}

func (fake *Converger) ConvergeCallCount() int {
	fake.convergeMutex.RLock()
	defer fake.convergeMutex.RUnlock()
	return len(fake.convergeArgsForCall)
}

func (fake *Converger) ConvergeArgsForCall(i int) []controller.Lease {
	fake.convergeMutex.RLock()
	defer fake.convergeMutex.RUnlock()
	return fake.convergeArgsForCall[i].arg1
}

func (fake *Converger) ConvergeReturns(result1 error) {
	fake.ConvergeStub = nil
	fake.convergeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Converger) ConvergeReturnsOnCall(i int, result1 error) {
	fake.ConvergeStub = nil
	if fake.convergeReturnsOnCall == nil {
		fake.convergeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.convergeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Converger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convergeMutex.RLock()
	defer fake.convergeMutex.RUnlock()
	return fake.invocations
}

func (fake *Converger) recordInvocation(key string, args []interface{}) {
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
