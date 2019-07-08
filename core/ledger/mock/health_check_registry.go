// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	sync "sync"

	healthz "github.com/justledger/fabric-lib-go/healthz"
	ledger "github.com/justledger/fabric/core/ledger"
)

type HealthCheckRegistry struct {
	RegisterCheckerStub        func(string, healthz.HealthChecker) error
	registerCheckerMutex       sync.RWMutex
	registerCheckerArgsForCall []struct {
		arg1 string
		arg2 healthz.HealthChecker
	}
	registerCheckerReturns struct {
		result1 error
	}
	registerCheckerReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *HealthCheckRegistry) RegisterChecker(arg1 string, arg2 healthz.HealthChecker) error {
	fake.registerCheckerMutex.Lock()
	ret, specificReturn := fake.registerCheckerReturnsOnCall[len(fake.registerCheckerArgsForCall)]
	fake.registerCheckerArgsForCall = append(fake.registerCheckerArgsForCall, struct {
		arg1 string
		arg2 healthz.HealthChecker
	}{arg1, arg2})
	fake.recordInvocation("RegisterChecker", []interface{}{arg1, arg2})
	fake.registerCheckerMutex.Unlock()
	if fake.RegisterCheckerStub != nil {
		return fake.RegisterCheckerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.registerCheckerReturns
	return fakeReturns.result1
}

func (fake *HealthCheckRegistry) RegisterCheckerCallCount() int {
	fake.registerCheckerMutex.RLock()
	defer fake.registerCheckerMutex.RUnlock()
	return len(fake.registerCheckerArgsForCall)
}

func (fake *HealthCheckRegistry) RegisterCheckerCalls(stub func(string, healthz.HealthChecker) error) {
	fake.registerCheckerMutex.Lock()
	defer fake.registerCheckerMutex.Unlock()
	fake.RegisterCheckerStub = stub
}

func (fake *HealthCheckRegistry) RegisterCheckerArgsForCall(i int) (string, healthz.HealthChecker) {
	fake.registerCheckerMutex.RLock()
	defer fake.registerCheckerMutex.RUnlock()
	argsForCall := fake.registerCheckerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *HealthCheckRegistry) RegisterCheckerReturns(result1 error) {
	fake.registerCheckerMutex.Lock()
	defer fake.registerCheckerMutex.Unlock()
	fake.RegisterCheckerStub = nil
	fake.registerCheckerReturns = struct {
		result1 error
	}{result1}
}

func (fake *HealthCheckRegistry) RegisterCheckerReturnsOnCall(i int, result1 error) {
	fake.registerCheckerMutex.Lock()
	defer fake.registerCheckerMutex.Unlock()
	fake.RegisterCheckerStub = nil
	if fake.registerCheckerReturnsOnCall == nil {
		fake.registerCheckerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerCheckerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *HealthCheckRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerCheckerMutex.RLock()
	defer fake.registerCheckerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *HealthCheckRegistry) recordInvocation(key string, args []interface{}) {
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

var _ ledger.HealthCheckRegistry = new(HealthCheckRegistry)
