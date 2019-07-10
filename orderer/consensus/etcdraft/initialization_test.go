/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package etcdraft_test

import (
	"testing"

	"justledgercommon/metrics/disabled"
	"justledgercore/comm"
	"justledgerorderer/common/cluster"
	"justledgerorderer/common/localconfig"
	"justledgerorderer/common/multichannel"
	"justledgerorderer/consensus/etcdraft"
	"justledgerorderer/consensus/etcdraft/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewEtcdRaftConsenter(t *testing.T) {
	srv, err := comm.NewGRPCServer("127.0.0.1:0", comm.ServerConfig{})
	assert.NoError(t, err)
	defer srv.Stop()
	dialer := &cluster.PredicateDialer{}
	consenter := etcdraft.New(dialer,
		&localconfig.TopLevel{},
		comm.ServerConfig{
			SecOpts: &comm.SecureOptions{
				Certificate: []byte{1, 2, 3},
			},
		}, srv, &multichannel.Registrar{},
		&mocks.InactiveChainRegistry{}, &disabled.Provider{})

	// Assert that the certificate from the gRPC server was passed to the consenter
	assert.Equal(t, []byte{1, 2, 3}, consenter.Cert)
	// Assert that all dependencies for the consenter were populated
	assert.NotNil(t, consenter.Communication)
	assert.NotNil(t, consenter.Chains)
	assert.NotNil(t, consenter.ChainSelector)
	assert.NotNil(t, consenter.Dispatcher)
	assert.NotNil(t, consenter.Logger)
}
