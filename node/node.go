package node

import (
	"log"
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
	proto.UnimplementedRicartAgrawalaServer

	id              string
	port            string
	clock           *clock.LamportClock
	state           NodeState
	replyCount      int
	deferredReplies []string

	logger *log.Logger

	mu sync.Mutex

	// Server
	grpcServer *grpc.Server

	// Clients
	peerConnections map[string]*grpc.ClientConn
	peers           map[string]proto.RicartAgrawalaClient

	peersConfig Config
}

type Config struct {
	Nodes []struct {
		Id   string `json:"id"`
		Port string `json:"port"`
	} `json:"nodes"`
}

func NewNode(index int, config *Config) (*Node, error) {
	n := config.Nodes[index]
	logger := log.New(os.Stderr, "", 0)

	node := &Node{
		id:              n.Id,
		port:            n.Port,
		clock:           &clock.LamportClock{},
		state:           Released,
		replyCount:      0,
		deferredReplies: make([]string, 0),

		logger: logger,

		grpcServer: grpc.NewServer(),

		peerConnections: make(map[string]*grpc.ClientConn),
		peers:           make(map[string]proto.RicartAgrawalaClient),
		peersConfig:     *config,
	}

	return node, nil
}

func (n *Node) Start() error {
	n.logger.Println("starting node")

	go n.startServer()
	n.connectToPeers()
	// n.simulate()

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
