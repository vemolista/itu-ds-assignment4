package node

import (
	"context"
	"time"

	"github.com/vemolista/itu-ds-assignment4/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (n *Node) connectToPeers() error {
	for _, peer := range n.peersConfig.Nodes {
		if peer.Id == n.id {
			// Do not connect to self
			n.logger.Printf("skipping connection to self, peer.Id: %s, self.Id: %s", peer.Id, n.id)
			continue
		}

		var conn *grpc.ClientConn
		var err error

		for {
			conn, err = grpc.NewClient("localhost"+peer.Port,
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			if err != nil {
				n.logger.Printf("failed to create connection to peer.Id: %s", peer.Id)
			}

			client := proto.NewRicartAgrawalaClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			_, err = client.HealthCheck(ctx, &proto.Empty{})
			if err == nil {
				cancel()
				break
			}
			cancel()

			n.logger.Printf("failed to connect to peer.Id: %s. retrying.\n", peer.Id)
			conn.Close()
			time.Sleep(500 * time.Millisecond)

		}

		n.logger.Printf("connected to peer: %s", peer.Id)

		n.peerConnections[peer.Id] = conn
		n.peers[peer.Id] = proto.NewRicartAgrawalaClient(conn)
	}

	return nil
}
