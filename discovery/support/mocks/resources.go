// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/justledger/fabric/common/channelconfig"
	"github.com/justledger/fabric/common/configtx"
	"github.com/justledger/fabric/common/policies"
	"github.com/justledger/fabric/msp"
)

type Resources struct {
	ConfigtxValidatorStub        func() configtx.Validator
	configtxValidatorMutex       sync.RWMutex
	configtxValidatorArgsForCall []struct{}
	configtxValidatorReturns     struct {
		result1 configtx.Validator
	}
	configtxValidatorReturnsOnCall map[int]struct {
		result1 configtx.Validator
	}
	PolicyManagerStub        func() policies.Manager
	policyManagerMutex       sync.RWMutex
	policyManagerArgsForCall []struct{}
	policyManagerReturns     struct {
		result1 policies.Manager
	}
	policyManagerReturnsOnCall map[int]struct {
		result1 policies.Manager
	}
	ChannelConfigStub        func() channelconfig.Channel
	channelConfigMutex       sync.RWMutex
	channelConfigArgsForCall []struct{}
	channelConfigReturns     struct {
		result1 channelconfig.Channel
	}
	channelConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Channel
	}
	OrdererConfigStub        func() (channelconfig.Orderer, bool)
	ordererConfigMutex       sync.RWMutex
	ordererConfigArgsForCall []struct{}
	ordererConfigReturns     struct {
		result1 channelconfig.Orderer
		result2 bool
	}
	ordererConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Orderer
		result2 bool
	}
	ConsortiumsConfigStub        func() (channelconfig.Consortiums, bool)
	consortiumsConfigMutex       sync.RWMutex
	consortiumsConfigArgsForCall []struct{}
	consortiumsConfigReturns     struct {
		result1 channelconfig.Consortiums
		result2 bool
	}
	consortiumsConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Consortiums
		result2 bool
	}
	ApplicationConfigStub        func() (channelconfig.Application, bool)
	applicationConfigMutex       sync.RWMutex
	applicationConfigArgsForCall []struct{}
	applicationConfigReturns     struct {
		result1 channelconfig.Application
		result2 bool
	}
	applicationConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Application
		result2 bool
	}
	MSPManagerStub        func() msp.MSPManager
	mSPManagerMutex       sync.RWMutex
	mSPManagerArgsForCall []struct{}
	mSPManagerReturns     struct {
		result1 msp.MSPManager
	}
	mSPManagerReturnsOnCall map[int]struct {
		result1 msp.MSPManager
	}
	ValidateNewStub        func(resources channelconfig.Resources) error
	validateNewMutex       sync.RWMutex
	validateNewArgsForCall []struct {
		resources channelconfig.Resources
	}
	validateNewReturns struct {
		result1 error
	}
	validateNewReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Resources) ConfigtxValidator() configtx.Validator {
	fake.configtxValidatorMutex.Lock()
	ret, specificReturn := fake.configtxValidatorReturnsOnCall[len(fake.configtxValidatorArgsForCall)]
	fake.configtxValidatorArgsForCall = append(fake.configtxValidatorArgsForCall, struct{}{})
	fake.recordInvocation("ConfigtxValidator", []interface{}{})
	fake.configtxValidatorMutex.Unlock()
	if fake.ConfigtxValidatorStub != nil {
		return fake.ConfigtxValidatorStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.configtxValidatorReturns.result1
}

func (fake *Resources) ConfigtxValidatorCallCount() int {
	fake.configtxValidatorMutex.RLock()
	defer fake.configtxValidatorMutex.RUnlock()
	return len(fake.configtxValidatorArgsForCall)
}

func (fake *Resources) ConfigtxValidatorReturns(result1 configtx.Validator) {
	fake.ConfigtxValidatorStub = nil
	fake.configtxValidatorReturns = struct {
		result1 configtx.Validator
	}{result1}
}

func (fake *Resources) ConfigtxValidatorReturnsOnCall(i int, result1 configtx.Validator) {
	fake.ConfigtxValidatorStub = nil
	if fake.configtxValidatorReturnsOnCall == nil {
		fake.configtxValidatorReturnsOnCall = make(map[int]struct {
			result1 configtx.Validator
		})
	}
	fake.configtxValidatorReturnsOnCall[i] = struct {
		result1 configtx.Validator
	}{result1}
}

func (fake *Resources) PolicyManager() policies.Manager {
	fake.policyManagerMutex.Lock()
	ret, specificReturn := fake.policyManagerReturnsOnCall[len(fake.policyManagerArgsForCall)]
	fake.policyManagerArgsForCall = append(fake.policyManagerArgsForCall, struct{}{})
	fake.recordInvocation("PolicyManager", []interface{}{})
	fake.policyManagerMutex.Unlock()
	if fake.PolicyManagerStub != nil {
		return fake.PolicyManagerStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.policyManagerReturns.result1
}

func (fake *Resources) PolicyManagerCallCount() int {
	fake.policyManagerMutex.RLock()
	defer fake.policyManagerMutex.RUnlock()
	return len(fake.policyManagerArgsForCall)
}

func (fake *Resources) PolicyManagerReturns(result1 policies.Manager) {
	fake.PolicyManagerStub = nil
	fake.policyManagerReturns = struct {
		result1 policies.Manager
	}{result1}
}

func (fake *Resources) PolicyManagerReturnsOnCall(i int, result1 policies.Manager) {
	fake.PolicyManagerStub = nil
	if fake.policyManagerReturnsOnCall == nil {
		fake.policyManagerReturnsOnCall = make(map[int]struct {
			result1 policies.Manager
		})
	}
	fake.policyManagerReturnsOnCall[i] = struct {
		result1 policies.Manager
	}{result1}
}

func (fake *Resources) ChannelConfig() channelconfig.Channel {
	fake.channelConfigMutex.Lock()
	ret, specificReturn := fake.channelConfigReturnsOnCall[len(fake.channelConfigArgsForCall)]
	fake.channelConfigArgsForCall = append(fake.channelConfigArgsForCall, struct{}{})
	fake.recordInvocation("ChannelConfig", []interface{}{})
	fake.channelConfigMutex.Unlock()
	if fake.ChannelConfigStub != nil {
		return fake.ChannelConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.channelConfigReturns.result1
}

func (fake *Resources) ChannelConfigCallCount() int {
	fake.channelConfigMutex.RLock()
	defer fake.channelConfigMutex.RUnlock()
	return len(fake.channelConfigArgsForCall)
}

func (fake *Resources) ChannelConfigReturns(result1 channelconfig.Channel) {
	fake.ChannelConfigStub = nil
	fake.channelConfigReturns = struct {
		result1 channelconfig.Channel
	}{result1}
}

func (fake *Resources) ChannelConfigReturnsOnCall(i int, result1 channelconfig.Channel) {
	fake.ChannelConfigStub = nil
	if fake.channelConfigReturnsOnCall == nil {
		fake.channelConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Channel
		})
	}
	fake.channelConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Channel
	}{result1}
}

func (fake *Resources) OrdererConfig() (channelconfig.Orderer, bool) {
	fake.ordererConfigMutex.Lock()
	ret, specificReturn := fake.ordererConfigReturnsOnCall[len(fake.ordererConfigArgsForCall)]
	fake.ordererConfigArgsForCall = append(fake.ordererConfigArgsForCall, struct{}{})
	fake.recordInvocation("OrdererConfig", []interface{}{})
	fake.ordererConfigMutex.Unlock()
	if fake.OrdererConfigStub != nil {
		return fake.OrdererConfigStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.ordererConfigReturns.result1, fake.ordererConfigReturns.result2
}

func (fake *Resources) OrdererConfigCallCount() int {
	fake.ordererConfigMutex.RLock()
	defer fake.ordererConfigMutex.RUnlock()
	return len(fake.ordererConfigArgsForCall)
}

func (fake *Resources) OrdererConfigReturns(result1 channelconfig.Orderer, result2 bool) {
	fake.OrdererConfigStub = nil
	fake.ordererConfigReturns = struct {
		result1 channelconfig.Orderer
		result2 bool
	}{result1, result2}
}

func (fake *Resources) OrdererConfigReturnsOnCall(i int, result1 channelconfig.Orderer, result2 bool) {
	fake.OrdererConfigStub = nil
	if fake.ordererConfigReturnsOnCall == nil {
		fake.ordererConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Orderer
			result2 bool
		})
	}
	fake.ordererConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Orderer
		result2 bool
	}{result1, result2}
}

func (fake *Resources) ConsortiumsConfig() (channelconfig.Consortiums, bool) {
	fake.consortiumsConfigMutex.Lock()
	ret, specificReturn := fake.consortiumsConfigReturnsOnCall[len(fake.consortiumsConfigArgsForCall)]
	fake.consortiumsConfigArgsForCall = append(fake.consortiumsConfigArgsForCall, struct{}{})
	fake.recordInvocation("ConsortiumsConfig", []interface{}{})
	fake.consortiumsConfigMutex.Unlock()
	if fake.ConsortiumsConfigStub != nil {
		return fake.ConsortiumsConfigStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.consortiumsConfigReturns.result1, fake.consortiumsConfigReturns.result2
}

func (fake *Resources) ConsortiumsConfigCallCount() int {
	fake.consortiumsConfigMutex.RLock()
	defer fake.consortiumsConfigMutex.RUnlock()
	return len(fake.consortiumsConfigArgsForCall)
}

func (fake *Resources) ConsortiumsConfigReturns(result1 channelconfig.Consortiums, result2 bool) {
	fake.ConsortiumsConfigStub = nil
	fake.consortiumsConfigReturns = struct {
		result1 channelconfig.Consortiums
		result2 bool
	}{result1, result2}
}

func (fake *Resources) ConsortiumsConfigReturnsOnCall(i int, result1 channelconfig.Consortiums, result2 bool) {
	fake.ConsortiumsConfigStub = nil
	if fake.consortiumsConfigReturnsOnCall == nil {
		fake.consortiumsConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Consortiums
			result2 bool
		})
	}
	fake.consortiumsConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Consortiums
		result2 bool
	}{result1, result2}
}

func (fake *Resources) ApplicationConfig() (channelconfig.Application, bool) {
	fake.applicationConfigMutex.Lock()
	ret, specificReturn := fake.applicationConfigReturnsOnCall[len(fake.applicationConfigArgsForCall)]
	fake.applicationConfigArgsForCall = append(fake.applicationConfigArgsForCall, struct{}{})
	fake.recordInvocation("ApplicationConfig", []interface{}{})
	fake.applicationConfigMutex.Unlock()
	if fake.ApplicationConfigStub != nil {
		return fake.ApplicationConfigStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.applicationConfigReturns.result1, fake.applicationConfigReturns.result2
}

func (fake *Resources) ApplicationConfigCallCount() int {
	fake.applicationConfigMutex.RLock()
	defer fake.applicationConfigMutex.RUnlock()
	return len(fake.applicationConfigArgsForCall)
}

func (fake *Resources) ApplicationConfigReturns(result1 channelconfig.Application, result2 bool) {
	fake.ApplicationConfigStub = nil
	fake.applicationConfigReturns = struct {
		result1 channelconfig.Application
		result2 bool
	}{result1, result2}
}

func (fake *Resources) ApplicationConfigReturnsOnCall(i int, result1 channelconfig.Application, result2 bool) {
	fake.ApplicationConfigStub = nil
	if fake.applicationConfigReturnsOnCall == nil {
		fake.applicationConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Application
			result2 bool
		})
	}
	fake.applicationConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Application
		result2 bool
	}{result1, result2}
}

func (fake *Resources) MSPManager() msp.MSPManager {
	fake.mSPManagerMutex.Lock()
	ret, specificReturn := fake.mSPManagerReturnsOnCall[len(fake.mSPManagerArgsForCall)]
	fake.mSPManagerArgsForCall = append(fake.mSPManagerArgsForCall, struct{}{})
	fake.recordInvocation("MSPManager", []interface{}{})
	fake.mSPManagerMutex.Unlock()
	if fake.MSPManagerStub != nil {
		return fake.MSPManagerStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.mSPManagerReturns.result1
}

func (fake *Resources) MSPManagerCallCount() int {
	fake.mSPManagerMutex.RLock()
	defer fake.mSPManagerMutex.RUnlock()
	return len(fake.mSPManagerArgsForCall)
}

func (fake *Resources) MSPManagerReturns(result1 msp.MSPManager) {
	fake.MSPManagerStub = nil
	fake.mSPManagerReturns = struct {
		result1 msp.MSPManager
	}{result1}
}

func (fake *Resources) MSPManagerReturnsOnCall(i int, result1 msp.MSPManager) {
	fake.MSPManagerStub = nil
	if fake.mSPManagerReturnsOnCall == nil {
		fake.mSPManagerReturnsOnCall = make(map[int]struct {
			result1 msp.MSPManager
		})
	}
	fake.mSPManagerReturnsOnCall[i] = struct {
		result1 msp.MSPManager
	}{result1}
}

func (fake *Resources) ValidateNew(resources channelconfig.Resources) error {
	fake.validateNewMutex.Lock()
	ret, specificReturn := fake.validateNewReturnsOnCall[len(fake.validateNewArgsForCall)]
	fake.validateNewArgsForCall = append(fake.validateNewArgsForCall, struct {
		resources channelconfig.Resources
	}{resources})
	fake.recordInvocation("ValidateNew", []interface{}{resources})
	fake.validateNewMutex.Unlock()
	if fake.ValidateNewStub != nil {
		return fake.ValidateNewStub(resources)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateNewReturns.result1
}

func (fake *Resources) ValidateNewCallCount() int {
	fake.validateNewMutex.RLock()
	defer fake.validateNewMutex.RUnlock()
	return len(fake.validateNewArgsForCall)
}

func (fake *Resources) ValidateNewArgsForCall(i int) channelconfig.Resources {
	fake.validateNewMutex.RLock()
	defer fake.validateNewMutex.RUnlock()
	return fake.validateNewArgsForCall[i].resources
}

func (fake *Resources) ValidateNewReturns(result1 error) {
	fake.ValidateNewStub = nil
	fake.validateNewReturns = struct {
		result1 error
	}{result1}
}

func (fake *Resources) ValidateNewReturnsOnCall(i int, result1 error) {
	fake.ValidateNewStub = nil
	if fake.validateNewReturnsOnCall == nil {
		fake.validateNewReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateNewReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Resources) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configtxValidatorMutex.RLock()
	defer fake.configtxValidatorMutex.RUnlock()
	fake.policyManagerMutex.RLock()
	defer fake.policyManagerMutex.RUnlock()
	fake.channelConfigMutex.RLock()
	defer fake.channelConfigMutex.RUnlock()
	fake.ordererConfigMutex.RLock()
	defer fake.ordererConfigMutex.RUnlock()
	fake.consortiumsConfigMutex.RLock()
	defer fake.consortiumsConfigMutex.RUnlock()
	fake.applicationConfigMutex.RLock()
	defer fake.applicationConfigMutex.RUnlock()
	fake.mSPManagerMutex.RLock()
	defer fake.mSPManagerMutex.RUnlock()
	fake.validateNewMutex.RLock()
	defer fake.validateNewMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Resources) recordInvocation(key string, args []interface{}) {
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

var _ channelconfig.Resources = new(Resources)
