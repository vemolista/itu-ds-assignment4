package node

import (
	"sync"

	"github.com/vemolista/itu-ds-assignment4/clock"
	"github.com/vemolista/itu-ds-assignment4/proto"

	"google.golang.org/grpc"
)

type NodeState int

const (
	Wanted NodeState = iota
	Released
	Held
)

type Node struct {
	id              string
	port            string
	clock           *clock.LamportClock
	requesting      bool
	replyCount      int
	deferredReplies []string

	// Server
	grpcServer *grpc.Server

	// Clients
	peers map[string]proto.RicartAgrawalaClient

	mu sync.Mutex
}

func NewNode(id, port, configPath string) (*Node, error) {
	// node constructor

	return nil, nil
}

func (n *Node) Start() error {
	// Start gRPC server
	// Connect to peers
	// Start simulation

	return nil
}

func (n *Node) RequestCriticalSection() {
	// Increment clock
	// Set state = wanted
	// Send RequestAccess to all peers
	// Wait for replies
}

func (n *Node) EnterCriticalSection() {
	// Log entry
	// Simulate work (sleep)
	// Log exit
}

func (n *Node) ReleaseCriticalSection() {
	// Send REPLY to deferred nodes
	// Set state to released
}
