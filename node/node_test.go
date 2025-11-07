package node

import (
	"testing"
)

func TestNodeConstructor(t *testing.T) {
	n, err := NewNode("node1", "localhost:50051", "../config.json")
	if err != nil {
		t.Fatalf("Failed to create node: %v", err)
	}

	if len(n.peersConfig.Nodes) != 5 {
		t.Errorf("Expected 5 configured nodes, got %d", len(n.peersConfig.Nodes))
	}
}
