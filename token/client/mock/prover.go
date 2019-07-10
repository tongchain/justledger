// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	sync "sync"

	token "justledgerprotos/token"
	tokena "justledgertoken"
	client "justledgertoken/client"
)

type Prover struct {
	RequestImportStub        func([]*token.TokenToIssue, tokena.SigningIdentity) ([]byte, error)
	requestImportMutex       sync.RWMutex
	requestImportArgsForCall []struct {
		arg1 []*token.TokenToIssue
		arg2 tokena.SigningIdentity
	}
	requestImportReturns struct {
		result1 []byte
		result2 error
	}
	requestImportReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	RequestTransferStub        func([][]byte, []*token.RecipientTransferShare, tokena.SigningIdentity) ([]byte, error)
	requestTransferMutex       sync.RWMutex
	requestTransferArgsForCall []struct {
		arg1 [][]byte
		arg2 []*token.RecipientTransferShare
		arg3 tokena.SigningIdentity
	}
	requestTransferReturns struct {
		result1 []byte
		result2 error
	}
	requestTransferReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Prover) RequestImport(arg1 []*token.TokenToIssue, arg2 tokena.SigningIdentity) ([]byte, error) {
	var arg1Copy []*token.TokenToIssue
	if arg1 != nil {
		arg1Copy = make([]*token.TokenToIssue, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.requestImportMutex.Lock()
	ret, specificReturn := fake.requestImportReturnsOnCall[len(fake.requestImportArgsForCall)]
	fake.requestImportArgsForCall = append(fake.requestImportArgsForCall, struct {
		arg1 []*token.TokenToIssue
		arg2 tokena.SigningIdentity
	}{arg1Copy, arg2})
	fake.recordInvocation("RequestImport", []interface{}{arg1Copy, arg2})
	fake.requestImportMutex.Unlock()
	if fake.RequestImportStub != nil {
		return fake.RequestImportStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.requestImportReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Prover) RequestImportCallCount() int {
	fake.requestImportMutex.RLock()
	defer fake.requestImportMutex.RUnlock()
	return len(fake.requestImportArgsForCall)
}

func (fake *Prover) RequestImportCalls(stub func([]*token.TokenToIssue, tokena.SigningIdentity) ([]byte, error)) {
	fake.requestImportMutex.Lock()
	defer fake.requestImportMutex.Unlock()
	fake.RequestImportStub = stub
}

func (fake *Prover) RequestImportArgsForCall(i int) ([]*token.TokenToIssue, tokena.SigningIdentity) {
	fake.requestImportMutex.RLock()
	defer fake.requestImportMutex.RUnlock()
	argsForCall := fake.requestImportArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Prover) RequestImportReturns(result1 []byte, result2 error) {
	fake.requestImportMutex.Lock()
	defer fake.requestImportMutex.Unlock()
	fake.RequestImportStub = nil
	fake.requestImportReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Prover) RequestImportReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.requestImportMutex.Lock()
	defer fake.requestImportMutex.Unlock()
	fake.RequestImportStub = nil
	if fake.requestImportReturnsOnCall == nil {
		fake.requestImportReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.requestImportReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Prover) RequestTransfer(arg1 [][]byte, arg2 []*token.RecipientTransferShare, arg3 tokena.SigningIdentity) ([]byte, error) {
	var arg1Copy [][]byte
	if arg1 != nil {
		arg1Copy = make([][]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []*token.RecipientTransferShare
	if arg2 != nil {
		arg2Copy = make([]*token.RecipientTransferShare, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.requestTransferMutex.Lock()
	ret, specificReturn := fake.requestTransferReturnsOnCall[len(fake.requestTransferArgsForCall)]
	fake.requestTransferArgsForCall = append(fake.requestTransferArgsForCall, struct {
		arg1 [][]byte
		arg2 []*token.RecipientTransferShare
		arg3 tokena.SigningIdentity
	}{arg1Copy, arg2Copy, arg3})
	fake.recordInvocation("RequestTransfer", []interface{}{arg1Copy, arg2Copy, arg3})
	fake.requestTransferMutex.Unlock()
	if fake.RequestTransferStub != nil {
		return fake.RequestTransferStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.requestTransferReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Prover) RequestTransferCallCount() int {
	fake.requestTransferMutex.RLock()
	defer fake.requestTransferMutex.RUnlock()
	return len(fake.requestTransferArgsForCall)
}

func (fake *Prover) RequestTransferCalls(stub func([][]byte, []*token.RecipientTransferShare, tokena.SigningIdentity) ([]byte, error)) {
	fake.requestTransferMutex.Lock()
	defer fake.requestTransferMutex.Unlock()
	fake.RequestTransferStub = stub
}

func (fake *Prover) RequestTransferArgsForCall(i int) ([][]byte, []*token.RecipientTransferShare, tokena.SigningIdentity) {
	fake.requestTransferMutex.RLock()
	defer fake.requestTransferMutex.RUnlock()
	argsForCall := fake.requestTransferArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *Prover) RequestTransferReturns(result1 []byte, result2 error) {
	fake.requestTransferMutex.Lock()
	defer fake.requestTransferMutex.Unlock()
	fake.RequestTransferStub = nil
	fake.requestTransferReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Prover) RequestTransferReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.requestTransferMutex.Lock()
	defer fake.requestTransferMutex.Unlock()
	fake.RequestTransferStub = nil
	if fake.requestTransferReturnsOnCall == nil {
		fake.requestTransferReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.requestTransferReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Prover) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.requestImportMutex.RLock()
	defer fake.requestImportMutex.RUnlock()
	fake.requestTransferMutex.RLock()
	defer fake.requestTransferMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Prover) recordInvocation(key string, args []interface{}) {
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

var _ client.Prover = new(Prover)
