/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package test

import (
	"justledger/fabric/common/channelconfig"
	"justledger/fabric/common/flogging"
	"justledger/fabric/common/genesis"
	"justledger/fabric/common/tools/configtxgen/configtxgentest"
	"justledger/fabric/common/tools/configtxgen/encoder"
	genesisconfig "justledger/fabric/common/tools/configtxgen/localconfig"
	"justledger/fabric/core/ledger/util"
	cb "justledger/fabric/protos/common"
	mspproto "justledger/fabric/protos/msp"
	"justledger/fabric/protos/peer"
	pb "justledger/fabric/protos/peer"
	"justledger/fabric/protos/utils"
)

var logger = flogging.MustGetLogger("common.configtx.test")

// MakeGenesisBlock creates a genesis block using the test templates for the given chainID
func MakeGenesisBlock(chainID string) (*cb.Block, error) {
	profile := configtxgentest.Load(genesisconfig.SampleDevModeSoloProfile)
	channelGroup, err := encoder.NewChannelGroup(profile)
	if err != nil {
		logger.Panicf("Error creating channel config: %s", err)
	}

	gb := genesis.NewFactoryImpl(channelGroup).Block(chainID)
	if gb == nil {
		return gb, nil
	}

	txsFilter := util.NewTxValidationFlagsSetValue(len(gb.Data.Data), peer.TxValidationCode_VALID)
	gb.Metadata.Metadata[cb.BlockMetadataIndex_TRANSACTIONS_FILTER] = txsFilter

	return gb, nil
}

// MakeGenesisBlockWithMSPs creates a genesis block using the MSPs provided for the given chainID
func MakeGenesisBlockFromMSPs(chainID string, appMSPConf, ordererMSPConf *mspproto.MSPConfig, appOrgID, ordererOrgID string) (*cb.Block, error) {
	profile := configtxgentest.Load(genesisconfig.SampleDevModeSoloProfile)
	profile.Orderer.Organizations = nil
	channelGroup, err := encoder.NewChannelGroup(profile)
	if err != nil {
		logger.Panicf("Error creating channel config: %s", err)
	}

	ordererOrg := cb.NewConfigGroup()
	ordererOrg.ModPolicy = channelconfig.AdminsPolicyKey
	ordererOrg.Values[channelconfig.MSPKey] = &cb.ConfigValue{
		Value:     utils.MarshalOrPanic(channelconfig.MSPValue(ordererMSPConf).Value()),
		ModPolicy: channelconfig.AdminsPolicyKey,
	}

	applicationOrg := cb.NewConfigGroup()
	applicationOrg.ModPolicy = channelconfig.AdminsPolicyKey
	applicationOrg.Values[channelconfig.MSPKey] = &cb.ConfigValue{
		Value:     utils.MarshalOrPanic(channelconfig.MSPValue(appMSPConf).Value()),
		ModPolicy: channelconfig.AdminsPolicyKey,
	}
	applicationOrg.Values[channelconfig.AnchorPeersKey] = &cb.ConfigValue{
		Value:     utils.MarshalOrPanic(channelconfig.AnchorPeersValue([]*pb.AnchorPeer{}).Value()),
		ModPolicy: channelconfig.AdminsPolicyKey,
	}

	channelGroup.Groups[channelconfig.OrdererGroupKey].Groups[ordererOrgID] = ordererOrg
	channelGroup.Groups[channelconfig.ApplicationGroupKey].Groups[appOrgID] = applicationOrg

	return genesis.NewFactoryImpl(channelGroup).Block(chainID), nil
}
