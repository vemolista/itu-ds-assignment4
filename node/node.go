package node

import (
	"context"
	"log"
	"math/rand/v2"
	"os"
	"sync"
	"time"

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
	deferredReplies map[string]chan struct{}

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
		deferredReplies: make(map[string]chan struct{}),

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

	// TODO: wait until all peers are connected

	n.simulate()

	return nil
}

func (n *Node) simulate() {
	for {
		time.Sleep(time.Duration(rand.IntN(5)+3) * time.Second)
		n.RequestCriticalSection()
		n.logger.Println("in critical section")

		// working in the critical section
		time.Sleep(time.Duration(rand.IntN(3000)) * time.Millisecond)

		n.ReleaseCriticalSection()
	}
}

func (n *Node) RequestCriticalSection() {
	// TODO: Increment clock

	n.mu.Lock()
	n.state = Wanted
	n.mu.Unlock()

	for _, peer := range n.peers {
		peer.RequestAccess(context.Background(), &proto.Request{
			Timestamp: n.clock.Get(),
			NodeId:    n.id,
		})
	}

	n.mu.Lock()
	n.state = Held
	n.mu.Unlock()
}

func (n *Node) ReleaseCriticalSection() {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.state = Released
	n.logger.Println("state set to released")

	for _, v := range n.deferredReplies {
		v <- struct{}{}
	}
}
