// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	persistence_test "justledger/core/chaincode/persistence"
)

type PackageParser struct {
	ParseStub        func(data []byte) (*persistence_test.ChaincodePackage, error)
	parseMutex       sync.RWMutex
	parseArgsForCall []struct {
		data []byte
	}
	parseReturns struct {
		result1 *persistence_test.ChaincodePackage
		result2 error
	}
	parseReturnsOnCall map[int]struct {
		result1 *persistence_test.ChaincodePackage
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PackageParser) Parse(data []byte) (*persistence_test.ChaincodePackage, error) {
	var dataCopy []byte
	if data != nil {
		dataCopy = make([]byte, len(data))
		copy(dataCopy, data)
	}
	fake.parseMutex.Lock()
	ret, specificReturn := fake.parseReturnsOnCall[len(fake.parseArgsForCall)]
	fake.parseArgsForCall = append(fake.parseArgsForCall, struct {
		data []byte
	}{dataCopy})
	fake.recordInvocation("Parse", []interface{}{dataCopy})
	fake.parseMutex.Unlock()
	if fake.ParseStub != nil {
		return fake.ParseStub(data)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.parseReturns.result1, fake.parseReturns.result2
}

func (fake *PackageParser) ParseCallCount() int {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return len(fake.parseArgsForCall)
}

func (fake *PackageParser) ParseArgsForCall(i int) []byte {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return fake.parseArgsForCall[i].data
}

func (fake *PackageParser) ParseReturns(result1 *persistence_test.ChaincodePackage, result2 error) {
	fake.ParseStub = nil
	fake.parseReturns = struct {
		result1 *persistence_test.ChaincodePackage
		result2 error
	}{result1, result2}
}

func (fake *PackageParser) ParseReturnsOnCall(i int, result1 *persistence_test.ChaincodePackage, result2 error) {
	fake.ParseStub = nil
	if fake.parseReturnsOnCall == nil {
		fake.parseReturnsOnCall = make(map[int]struct {
			result1 *persistence_test.ChaincodePackage
			result2 error
		})
	}
	fake.parseReturnsOnCall[i] = struct {
		result1 *persistence_test.ChaincodePackage
		result2 error
	}{result1, result2}
}

func (fake *PackageParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PackageParser) recordInvocation(key string, args []interface{}) {
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
