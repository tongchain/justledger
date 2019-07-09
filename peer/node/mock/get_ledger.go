// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"justledger/fabric/core/ledger"
)

type GetLedger struct {
	Stub        func(string) ledger.PeerLedger
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 string
	}
	returns struct {
		result1 ledger.PeerLedger
	}
	returnsOnCall map[int]struct {
		result1 ledger.PeerLedger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *GetLedger) Spy(arg1 string) ledger.PeerLedger {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("getLedger", []interface{}{arg1})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.returns.result1
}

func (fake *GetLedger) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *GetLedger) Calls(stub func(string) ledger.PeerLedger) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *GetLedger) ArgsForCall(i int) string {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1
}

func (fake *GetLedger) Returns(result1 ledger.PeerLedger) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 ledger.PeerLedger
	}{result1}
}

func (fake *GetLedger) ReturnsOnCall(i int, result1 ledger.PeerLedger) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 ledger.PeerLedger
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 ledger.PeerLedger
	}{result1}
}

func (fake *GetLedger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *GetLedger) recordInvocation(key string, args []interface{}) {
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
