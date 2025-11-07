package node

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	state           NodeState
	replyCount      int
	deferredReplies []string

	// Server
	grpcServer *grpc.Server

	// Clients
	peerClients map[string]proto.RicartAgrawalaClient
	peersConfig Config

	mu sync.Mutex
}

type Config struct {
	Nodes []struct {
		Id      string `json:"id"`
		Address string `json:"address"`
	} `json:"nodes"`
}

func NewNode(id, port, configPath string) (*Node, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer configFile.Close()

	var config Config
	bytes, err := io.ReadAll(configFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read bytes: %w", err)
	}
	json.Unmarshal(bytes, &config)

	node := &Node{
		id:              id,
		port:            port,
		clock:           &clock.LamportClock{},
		state:           Released,
		replyCount:      0,
		deferredReplies: make([]string, 0),

		grpcServer: grpc.NewServer(),

		peerClients: make(map[string]proto.RicartAgrawalaClient),
		peersConfig: config,
	}

	return node, nil
}

func (n *Node) Start() error {
	n.startServer()
	n.connectToPeers()
	n.simulate()

	// Start gRPC server
	// Connect to peers
	// Start simulation

	return nil
}

func (n *Node) simulate() {

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
