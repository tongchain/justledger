// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import cmd "justledgerdiscovery/cmd"
import common "justledgercmd/common"
import discovery "justledgerdiscovery/client"
import mock "github.com/stretchr/testify/mock"

// Stub is an autogenerated mock type for the Stub type
type Stub struct {
	mock.Mock
}

// Send provides a mock function with given fields: server, conf, req
func (_m *Stub) Send(server string, conf common.Config, req *discovery.Request) (cmd.ServiceResponse, error) {
	ret := _m.Called(server, conf, req)

	var r0 cmd.ServiceResponse
	if rf, ok := ret.Get(0).(func(string, common.Config, *discovery.Request) cmd.ServiceResponse); ok {
		r0 = rf(server, conf, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cmd.ServiceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, common.Config, *discovery.Request) error); ok {
		r1 = rf(server, conf, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
