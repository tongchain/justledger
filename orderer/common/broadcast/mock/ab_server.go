// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/justledger/fabric/protos/common"
	"github.com/justledger/fabric/protos/orderer"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type ABServer struct {
	SendStub        func(*orderer.BroadcastResponse) error
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		arg1 *orderer.BroadcastResponse
	}
	sendReturns struct {
		result1 error
	}
	sendReturnsOnCall map[int]struct {
		result1 error
	}
	RecvStub        func() (*common.Envelope, error)
	recvMutex       sync.RWMutex
	recvArgsForCall []struct{}
	recvReturns     struct {
		result1 *common.Envelope
		result2 error
	}
	recvReturnsOnCall map[int]struct {
		result1 *common.Envelope
		result2 error
	}
	SetHeaderStub        func(metadata.MD) error
	setHeaderMutex       sync.RWMutex
	setHeaderArgsForCall []struct {
		arg1 metadata.MD
	}
	setHeaderReturns struct {
		result1 error
	}
	setHeaderReturnsOnCall map[int]struct {
		result1 error
	}
	SendHeaderStub        func(metadata.MD) error
	sendHeaderMutex       sync.RWMutex
	sendHeaderArgsForCall []struct {
		arg1 metadata.MD
	}
	sendHeaderReturns struct {
		result1 error
	}
	sendHeaderReturnsOnCall map[int]struct {
		result1 error
	}
	SetTrailerStub        func(metadata.MD)
	setTrailerMutex       sync.RWMutex
	setTrailerArgsForCall []struct {
		arg1 metadata.MD
	}
	ContextStub        func() context.Context
	contextMutex       sync.RWMutex
	contextArgsForCall []struct{}
	contextReturns     struct {
		result1 context.Context
	}
	contextReturnsOnCall map[int]struct {
		result1 context.Context
	}
	SendMsgStub        func(m interface{}) error
	sendMsgMutex       sync.RWMutex
	sendMsgArgsForCall []struct {
		m interface{}
	}
	sendMsgReturns struct {
		result1 error
	}
	sendMsgReturnsOnCall map[int]struct {
		result1 error
	}
	RecvMsgStub        func(m interface{}) error
	recvMsgMutex       sync.RWMutex
	recvMsgArgsForCall []struct {
		m interface{}
	}
	recvMsgReturns struct {
		result1 error
	}
	recvMsgReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ABServer) Send(arg1 *orderer.BroadcastResponse) error {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		arg1 *orderer.BroadcastResponse
	}{arg1})
	fake.recordInvocation("Send", []interface{}{arg1})
	fake.sendMutex.Unlock()
	if fake.SendStub != nil {
		return fake.SendStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendReturns.result1
}

func (fake *ABServer) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *ABServer) SendArgsForCall(i int) *orderer.BroadcastResponse {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return fake.sendArgsForCall[i].arg1
}

func (fake *ABServer) SendReturns(result1 error) {
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) SendReturnsOnCall(i int, result1 error) {
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) Recv() (*common.Envelope, error) {
	fake.recvMutex.Lock()
	ret, specificReturn := fake.recvReturnsOnCall[len(fake.recvArgsForCall)]
	fake.recvArgsForCall = append(fake.recvArgsForCall, struct{}{})
	fake.recordInvocation("Recv", []interface{}{})
	fake.recvMutex.Unlock()
	if fake.RecvStub != nil {
		return fake.RecvStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.recvReturns.result1, fake.recvReturns.result2
}

func (fake *ABServer) RecvCallCount() int {
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	return len(fake.recvArgsForCall)
}

func (fake *ABServer) RecvReturns(result1 *common.Envelope, result2 error) {
	fake.RecvStub = nil
	fake.recvReturns = struct {
		result1 *common.Envelope
		result2 error
	}{result1, result2}
}

func (fake *ABServer) RecvReturnsOnCall(i int, result1 *common.Envelope, result2 error) {
	fake.RecvStub = nil
	if fake.recvReturnsOnCall == nil {
		fake.recvReturnsOnCall = make(map[int]struct {
			result1 *common.Envelope
			result2 error
		})
	}
	fake.recvReturnsOnCall[i] = struct {
		result1 *common.Envelope
		result2 error
	}{result1, result2}
}

func (fake *ABServer) SetHeader(arg1 metadata.MD) error {
	fake.setHeaderMutex.Lock()
	ret, specificReturn := fake.setHeaderReturnsOnCall[len(fake.setHeaderArgsForCall)]
	fake.setHeaderArgsForCall = append(fake.setHeaderArgsForCall, struct {
		arg1 metadata.MD
	}{arg1})
	fake.recordInvocation("SetHeader", []interface{}{arg1})
	fake.setHeaderMutex.Unlock()
	if fake.SetHeaderStub != nil {
		return fake.SetHeaderStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setHeaderReturns.result1
}

func (fake *ABServer) SetHeaderCallCount() int {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	return len(fake.setHeaderArgsForCall)
}

func (fake *ABServer) SetHeaderArgsForCall(i int) metadata.MD {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	return fake.setHeaderArgsForCall[i].arg1
}

func (fake *ABServer) SetHeaderReturns(result1 error) {
	fake.SetHeaderStub = nil
	fake.setHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) SetHeaderReturnsOnCall(i int, result1 error) {
	fake.SetHeaderStub = nil
	if fake.setHeaderReturnsOnCall == nil {
		fake.setHeaderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setHeaderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) SendHeader(arg1 metadata.MD) error {
	fake.sendHeaderMutex.Lock()
	ret, specificReturn := fake.sendHeaderReturnsOnCall[len(fake.sendHeaderArgsForCall)]
	fake.sendHeaderArgsForCall = append(fake.sendHeaderArgsForCall, struct {
		arg1 metadata.MD
	}{arg1})
	fake.recordInvocation("SendHeader", []interface{}{arg1})
	fake.sendHeaderMutex.Unlock()
	if fake.SendHeaderStub != nil {
		return fake.SendHeaderStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendHeaderReturns.result1
}

func (fake *ABServer) SendHeaderCallCount() int {
	fake.sendHeaderMutex.RLock()
	defer fake.sendHeaderMutex.RUnlock()
	return len(fake.sendHeaderArgsForCall)
}

func (fake *ABServer) SendHeaderArgsForCall(i int) metadata.MD {
	fake.sendHeaderMutex.RLock()
	defer fake.sendHeaderMutex.RUnlock()
	return fake.sendHeaderArgsForCall[i].arg1
}

func (fake *ABServer) SendHeaderReturns(result1 error) {
	fake.SendHeaderStub = nil
	fake.sendHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) SendHeaderReturnsOnCall(i int, result1 error) {
	fake.SendHeaderStub = nil
	if fake.sendHeaderReturnsOnCall == nil {
		fake.sendHeaderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendHeaderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) SetTrailer(arg1 metadata.MD) {
	fake.setTrailerMutex.Lock()
	fake.setTrailerArgsForCall = append(fake.setTrailerArgsForCall, struct {
		arg1 metadata.MD
	}{arg1})
	fake.recordInvocation("SetTrailer", []interface{}{arg1})
	fake.setTrailerMutex.Unlock()
	if fake.SetTrailerStub != nil {
		fake.SetTrailerStub(arg1)
	}
}

func (fake *ABServer) SetTrailerCallCount() int {
	fake.setTrailerMutex.RLock()
	defer fake.setTrailerMutex.RUnlock()
	return len(fake.setTrailerArgsForCall)
}

func (fake *ABServer) SetTrailerArgsForCall(i int) metadata.MD {
	fake.setTrailerMutex.RLock()
	defer fake.setTrailerMutex.RUnlock()
	return fake.setTrailerArgsForCall[i].arg1
}

func (fake *ABServer) Context() context.Context {
	fake.contextMutex.Lock()
	ret, specificReturn := fake.contextReturnsOnCall[len(fake.contextArgsForCall)]
	fake.contextArgsForCall = append(fake.contextArgsForCall, struct{}{})
	fake.recordInvocation("Context", []interface{}{})
	fake.contextMutex.Unlock()
	if fake.ContextStub != nil {
		return fake.ContextStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.contextReturns.result1
}

func (fake *ABServer) ContextCallCount() int {
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return len(fake.contextArgsForCall)
}

func (fake *ABServer) ContextReturns(result1 context.Context) {
	fake.ContextStub = nil
	fake.contextReturns = struct {
		result1 context.Context
	}{result1}
}

func (fake *ABServer) ContextReturnsOnCall(i int, result1 context.Context) {
	fake.ContextStub = nil
	if fake.contextReturnsOnCall == nil {
		fake.contextReturnsOnCall = make(map[int]struct {
			result1 context.Context
		})
	}
	fake.contextReturnsOnCall[i] = struct {
		result1 context.Context
	}{result1}
}

func (fake *ABServer) SendMsg(m interface{}) error {
	fake.sendMsgMutex.Lock()
	ret, specificReturn := fake.sendMsgReturnsOnCall[len(fake.sendMsgArgsForCall)]
	fake.sendMsgArgsForCall = append(fake.sendMsgArgsForCall, struct {
		m interface{}
	}{m})
	fake.recordInvocation("SendMsg", []interface{}{m})
	fake.sendMsgMutex.Unlock()
	if fake.SendMsgStub != nil {
		return fake.SendMsgStub(m)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendMsgReturns.result1
}

func (fake *ABServer) SendMsgCallCount() int {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	return len(fake.sendMsgArgsForCall)
}

func (fake *ABServer) SendMsgArgsForCall(i int) interface{} {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	return fake.sendMsgArgsForCall[i].m
}

func (fake *ABServer) SendMsgReturns(result1 error) {
	fake.SendMsgStub = nil
	fake.sendMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) SendMsgReturnsOnCall(i int, result1 error) {
	fake.SendMsgStub = nil
	if fake.sendMsgReturnsOnCall == nil {
		fake.sendMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) RecvMsg(m interface{}) error {
	fake.recvMsgMutex.Lock()
	ret, specificReturn := fake.recvMsgReturnsOnCall[len(fake.recvMsgArgsForCall)]
	fake.recvMsgArgsForCall = append(fake.recvMsgArgsForCall, struct {
		m interface{}
	}{m})
	fake.recordInvocation("RecvMsg", []interface{}{m})
	fake.recvMsgMutex.Unlock()
	if fake.RecvMsgStub != nil {
		return fake.RecvMsgStub(m)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.recvMsgReturns.result1
}

func (fake *ABServer) RecvMsgCallCount() int {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	return len(fake.recvMsgArgsForCall)
}

func (fake *ABServer) RecvMsgArgsForCall(i int) interface{} {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	return fake.recvMsgArgsForCall[i].m
}

func (fake *ABServer) RecvMsgReturns(result1 error) {
	fake.RecvMsgStub = nil
	fake.recvMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) RecvMsgReturnsOnCall(i int, result1 error) {
	fake.RecvMsgStub = nil
	if fake.recvMsgReturnsOnCall == nil {
		fake.recvMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.recvMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ABServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	fake.sendHeaderMutex.RLock()
	defer fake.sendHeaderMutex.RUnlock()
	fake.setTrailerMutex.RLock()
	defer fake.setTrailerMutex.RUnlock()
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ABServer) recordInvocation(key string, args []interface{}) {
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
