package node

import (
	"context"

	"github.com/vemolista/itu-ds-assignment4/proto"
)

func (n *Node) RequestAccess(ctx context.Context, r *proto.RequestAccessRequest) (*proto.RequestAccessResponse, error) {
	// update clock

	// reply or wait

	return nil, nil
}

func (n *Node) Reply(ctx context.Context, r *proto.ReplyRequest) (*proto.ReplyResponse, error) {
	return nil, nil
}

func (n *Node) startServer() error {
	return nil
}
