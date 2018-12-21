/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package scc

import (
	//import system chaincodes here
	"justledger/core/aclmgmt"
	"justledger/core/common/ccprovider"
	"justledger/core/scc/cscc"
	"justledger/core/scc/lscc"
	"justledger/core/scc/qscc"
)

func builtInSystemChaincodes(ccp ccprovider.ChaincodeProvider, p *Provider, aclProvider aclmgmt.ACLProvider) []*SystemChaincode {
	return []*SystemChaincode{
		{
			Enabled:           true,
			Name:              "cscc",
			Path:              "justledger/core/scc/cscc",
			InitArgs:          nil,
			Chaincode:         cscc.New(ccp, p, aclProvider),
			InvokableExternal: true, // cscc is invoked to join a channel
		},
		{
			Enabled:           true,
			Name:              "lscc",
			Path:              "justledger/core/scc/lscc",
			InitArgs:          nil,
			Chaincode:         lscc.New(p, aclProvider),
			InvokableExternal: true, // lscc is invoked to deploy new chaincodes
			InvokableCC2CC:    true, // lscc can be invoked by other chaincodes
		},
		{
			Enabled:           true,
			Name:              "qscc",
			Path:              "justledger/core/scc/qscc",
			InitArgs:          nil,
			Chaincode:         qscc.New(aclProvider),
			InvokableExternal: true, // qscc can be invoked to retrieve blocks
			InvokableCC2CC:    true, // qscc can be invoked to retrieve blocks also by a cc
		},
	}
}

//DeploySysCCs is the hook for system chaincodes where system chaincodes are registered with the fabric
//note the chaincode must still be deployed and launched like a user chaincode will be
func (p *Provider) DeploySysCCs(chainID string, ccp ccprovider.ChaincodeProvider) {
	for _, sysCC := range p.SysCCs {
		sysCC.deploySysCC(chainID, ccp)
	}
}

//DeDeploySysCCs is used in unit tests to stop and remove the system chaincodes before
//restarting them in the same process. This allows clean start of the system
//in the same process
func (p *Provider) DeDeploySysCCs(chainID string, ccp ccprovider.ChaincodeProvider) {
	for _, sysCC := range p.SysCCs {
		sysCC.deDeploySysCC(chainID, ccp)
	}
}
