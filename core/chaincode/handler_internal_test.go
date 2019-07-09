/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chaincode

import (
	"justledger/fabric/core/common/sysccprovider"
	"justledger/fabric/core/container/ccintf"
	pb "justledger/fabric/protos/peer"
)

// Helpers to access unexported state.

func SetHandlerChaincodeID(h *Handler, chaincodeID *pb.ChaincodeID) {
	h.chaincodeID = chaincodeID
}

func SetHandlerChatStream(h *Handler, chatStream ccintf.ChaincodeStream) {
	h.chatStream = chatStream
}

func SetHandlerCCInstance(h *Handler, ccInstance *sysccprovider.ChaincodeInstance) {
	h.ccInstance = ccInstance
}
