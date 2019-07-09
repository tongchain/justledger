// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	sync "sync"

	client "justledger/fabric/token/client"
)

type FabricTxSubmitter struct {
	SubmitStub        func([]byte) error
	submitMutex       sync.RWMutex
	submitArgsForCall []struct {
		arg1 []byte
	}
	submitReturns struct {
		result1 error
	}
	submitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FabricTxSubmitter) Submit(arg1 []byte) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.submitMutex.Lock()
	ret, specificReturn := fake.submitReturnsOnCall[len(fake.submitArgsForCall)]
	fake.submitArgsForCall = append(fake.submitArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Submit", []interface{}{arg1Copy})
	fake.submitMutex.Unlock()
	if fake.SubmitStub != nil {
		return fake.SubmitStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.submitReturns
	return fakeReturns.result1
}

func (fake *FabricTxSubmitter) SubmitCallCount() int {
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	return len(fake.submitArgsForCall)
}

func (fake *FabricTxSubmitter) SubmitCalls(stub func([]byte) error) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = stub
}

func (fake *FabricTxSubmitter) SubmitArgsForCall(i int) []byte {
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	argsForCall := fake.submitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FabricTxSubmitter) SubmitReturns(result1 error) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = nil
	fake.submitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FabricTxSubmitter) SubmitReturnsOnCall(i int, result1 error) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = nil
	if fake.submitReturnsOnCall == nil {
		fake.submitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.submitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FabricTxSubmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FabricTxSubmitter) recordInvocation(key string, args []interface{}) {
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

var _ client.FabricTxSubmitter = new(FabricTxSubmitter)
