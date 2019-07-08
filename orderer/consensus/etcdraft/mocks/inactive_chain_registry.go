// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import common "github.com/justledger/fabric/protos/common"
import etcdraft "github.com/justledger/fabric/orderer/consensus/etcdraft"
import mock "github.com/stretchr/testify/mock"

// InactiveChainRegistry is an autogenerated mock type for the InactiveChainRegistry type
type InactiveChainRegistry struct {
	mock.Mock
}

// TrackChain provides a mock function with given fields: chainName, genesisBlock, createChain
func (_m *InactiveChainRegistry) TrackChain(chainName string, genesisBlock *common.Block, createChain etcdraft.CreateChainCallback) {
	_m.Called(chainName, genesisBlock, createChain)
}