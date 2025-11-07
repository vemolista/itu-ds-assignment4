package node

import (
	"context"
	"net"

	"github.com/vemolista/itu-ds-assignment4/proto"
)

func requesterIsSmaller(requesterClock, responderClock int64, requesterId, responderId string) bool {
	if requesterClock < responderClock {
		return true
	}

	if requesterClock > responderClock {
		return false
	}

	// Clocks must be equal at this point, compare ids

	if requesterId < responderId {
		return true
	}

	return false
}

func (n *Node) RequestAccess(ctx context.Context, r *proto.Request) (*proto.Reply, error) {
	n.mu.Lock()
	if (n.state == Held) || (n.state == Wanted && requesterIsSmaller(r.Timestamp, n.clock.Get(), r.NodeId, n.id)) {
		n.deferredReplies[r.NodeId] = make(chan struct{})
		n.mu.Unlock()

		<-n.deferredReplies[r.NodeId]

		n.mu.Lock()
		delete(n.deferredReplies, r.NodeId)
	}
	n.mu.Unlock()

	return &proto.Reply{}, nil
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
