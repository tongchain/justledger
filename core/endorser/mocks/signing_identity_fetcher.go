/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by mockery v1.0.0
package mocks

import endorsement "github.com/justledger/fabric/core/handlers/endorsement/api/identities"
import mock "github.com/stretchr/testify/mock"
import peer "github.com/justledger/fabric/protos/peer"

// SigningIdentityFetcher is an autogenerated mock type for the SigningIdentityFetcher type
type SigningIdentityFetcher struct {
	mock.Mock
}

// SigningIdentityForRequest provides a mock function with given fields: _a0
func (_m *SigningIdentityFetcher) SigningIdentityForRequest(_a0 *peer.SignedProposal) (endorsement.SigningIdentity, error) {
	ret := _m.Called(_a0)

	var r0 endorsement.SigningIdentity
	if rf, ok := ret.Get(0).(func(*peer.SignedProposal) endorsement.SigningIdentity); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(endorsement.SigningIdentity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*peer.SignedProposal) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
