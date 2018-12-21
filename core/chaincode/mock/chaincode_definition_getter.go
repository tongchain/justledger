// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"justledger/core/common/ccprovider"
	"justledger/core/ledger"
)

type ChaincodeDefinitionGetter struct {
	ChaincodeDefinitionStub        func(chaincodeName string, txSim ledger.QueryExecutor) (ccprovider.ChaincodeDefinition, error)
	chaincodeDefinitionMutex       sync.RWMutex
	chaincodeDefinitionArgsForCall []struct {
		chaincodeName string
		txSim         ledger.QueryExecutor
	}
	chaincodeDefinitionReturns struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}
	chaincodeDefinitionReturnsOnCall map[int]struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChaincodeDefinitionGetter) ChaincodeDefinition(chaincodeName string, txSim ledger.QueryExecutor) (ccprovider.ChaincodeDefinition, error) {
	fake.chaincodeDefinitionMutex.Lock()
	ret, specificReturn := fake.chaincodeDefinitionReturnsOnCall[len(fake.chaincodeDefinitionArgsForCall)]
	fake.chaincodeDefinitionArgsForCall = append(fake.chaincodeDefinitionArgsForCall, struct {
		chaincodeName string
		txSim         ledger.QueryExecutor
	}{chaincodeName, txSim})
	fake.recordInvocation("ChaincodeDefinition", []interface{}{chaincodeName, txSim})
	fake.chaincodeDefinitionMutex.Unlock()
	if fake.ChaincodeDefinitionStub != nil {
		return fake.ChaincodeDefinitionStub(chaincodeName, txSim)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.chaincodeDefinitionReturns.result1, fake.chaincodeDefinitionReturns.result2
}

func (fake *ChaincodeDefinitionGetter) ChaincodeDefinitionCallCount() int {
	fake.chaincodeDefinitionMutex.RLock()
	defer fake.chaincodeDefinitionMutex.RUnlock()
	return len(fake.chaincodeDefinitionArgsForCall)
}

func (fake *ChaincodeDefinitionGetter) ChaincodeDefinitionArgsForCall(i int) (string, ledger.QueryExecutor) {
	fake.chaincodeDefinitionMutex.RLock()
	defer fake.chaincodeDefinitionMutex.RUnlock()
	return fake.chaincodeDefinitionArgsForCall[i].chaincodeName, fake.chaincodeDefinitionArgsForCall[i].txSim
}

func (fake *ChaincodeDefinitionGetter) ChaincodeDefinitionReturns(result1 ccprovider.ChaincodeDefinition, result2 error) {
	fake.ChaincodeDefinitionStub = nil
	fake.chaincodeDefinitionReturns = struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeDefinitionGetter) ChaincodeDefinitionReturnsOnCall(i int, result1 ccprovider.ChaincodeDefinition, result2 error) {
	fake.ChaincodeDefinitionStub = nil
	if fake.chaincodeDefinitionReturnsOnCall == nil {
		fake.chaincodeDefinitionReturnsOnCall = make(map[int]struct {
			result1 ccprovider.ChaincodeDefinition
			result2 error
		})
	}
	fake.chaincodeDefinitionReturnsOnCall[i] = struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeDefinitionGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chaincodeDefinitionMutex.RLock()
	defer fake.chaincodeDefinitionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChaincodeDefinitionGetter) recordInvocation(key string, args []interface{}) {
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
