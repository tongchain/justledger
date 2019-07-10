/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package privacyenabledstate_test

import (
	"testing"

	"justledgercore/ledger/kvledger/txmgmt/privacyenabledstate"
	"justledgercore/ledger/kvledger/txmgmt/statedb/statecouchdb"
	"justledgercore/ledger/kvledger/txmgmt/statedb/stateleveldb"
	"justledgercore/ledger/mock"
	. "github.com/onsi/gomega"
)

func TestHealthCheckRegister(t *testing.T) {
	gt := NewGomegaWithT(t)
	fakeHealthCheckRegistry := &mock.HealthCheckRegistry{}

	dbProvider := &privacyenabledstate.CommonStorageDBProvider{
		VersionedDBProvider: &stateleveldb.VersionedDBProvider{},
		HealthCheckRegistry: fakeHealthCheckRegistry,
	}

	err := dbProvider.RegisterHealthChecker()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(fakeHealthCheckRegistry.RegisterCheckerCallCount()).To(Equal(0))

	dbProvider.VersionedDBProvider = &statecouchdb.VersionedDBProvider{}
	err = dbProvider.RegisterHealthChecker()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(fakeHealthCheckRegistry.RegisterCheckerCallCount()).To(Equal(1))

	arg1, arg2 := fakeHealthCheckRegistry.RegisterCheckerArgsForCall(0)
	gt.Expect(arg1).To(Equal("couchdb"))
	gt.Expect(arg2).NotTo(Equal(nil))
}
