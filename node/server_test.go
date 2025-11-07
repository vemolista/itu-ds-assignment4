package node

import (
	"fmt"
	"testing"
)

func TestRequesterIsSmaller(t *testing.T) {
	var comparisonTests = []struct {
		requesterClock int64
		responderClock int64
		requesterId    string
		responderId    string
		expected       bool
	}{
		{
			requesterClock: 1,
			responderClock: 2,
			expected:       true,
		},
		{
			requesterClock: 2,
			responderClock: 1,
			expected:       false,
		},
		{
			requesterClock: 1,
			responderClock: 1,
			requesterId:    "node1",
			responderId:    "node2",
			expected:       true,
		},
		{
			requesterClock: 1,
			responderClock: 1,
			requesterId:    "node2",
			responderId:    "node1",
			expected:       false,
		},
		{
			requesterClock: 1,
			responderClock: 1,
			requesterId:    "node1",
			responderId:    "node1",
			expected:       false,
		},
	}

	for _, tt := range comparisonTests {
		t.Run(fmt.Sprintf("%#v", tt), func(t *testing.T) {
			actual := requesterIsSmaller(tt.requesterClock, tt.responderClock, tt.requesterId, tt.responderId)
			if actual != tt.expected {
				t.Errorf("got '%v', expected '%v'", actual, tt.expected)
			}
		})
	}
}
