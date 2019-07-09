/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"context"

	"justledger/peer/chaincode/api"
	cb "justledger/protos/common"
	ab "justledger/protos/orderer"
	"google.golang.org/grpc"
)

//go:generate counterfeiter -o ../mock/deliverclient.go -fake-name DeliverClient . DeliverClient

// DeliverClient defines the interface for a deliver client
type DeliverClient interface {
	Deliver(ctx context.Context, opts ...grpc.CallOption) (DeliverService, error)
}

//go:generate counterfeiter -o ../mock/deliverservice.go -fake-name DeliverService . DeliverService

// DeliverService defines the interface for delivering blocks
type DeliverService interface {
	Send(*cb.Envelope) error
	Recv() (*ab.DeliverResponse, error)
	CloseSend() error
}

//go:generate counterfeiter -o ../mock/peerdeliverclient.go -fake-name PeerDeliverClient . PeerDeliverClient

// PeerDeliverClient defines the interface for a peer deliver client
type PeerDeliverClient interface {
	Deliver(ctx context.Context, opts ...grpc.CallOption) (api.Deliver, error)
	DeliverFiltered(ctx context.Context, opts ...grpc.CallOption) (api.Deliver, error)
}
