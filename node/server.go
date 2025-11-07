package node

import (
	"context"
	"net"

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

func (n *Node) HealthCheck(ctx context.Context, r *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (n *Node) startServer() error {
	n.logger.Println("server starting")
	proto.RegisterRicartAgrawalaServer(n.grpcServer, n)

	listener, err := net.Listen("tcp", n.port)
	if err != nil {
		n.logger.Fatalf("failed to create a %s listener on port %s: %v\n", "tcp", n.port, err)
	}

	if err := n.grpcServer.Serve(listener); err != nil {
		n.logger.Fatalf("failed to start serving requests: %v", err)
	}

	return nil
}
