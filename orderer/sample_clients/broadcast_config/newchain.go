// Copyright IBM Corp. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"justledger/fabric/common/localmsp"
	"justledger/fabric/common/tools/configtxgen/encoder"
	genesisconfig "justledger/fabric/common/tools/configtxgen/localconfig"
	cb "justledger/fabric/protos/common"
)

func newChainRequest(consensusType, creationPolicy, newChannelID string) *cb.Envelope {
	env, err := encoder.MakeChannelCreationTransaction(newChannelID, localmsp.NewSigner(), genesisconfig.Load(genesisconfig.SampleSingleMSPChannelProfile))
	if err != nil {
		panic(err)
	}
	return env
}
