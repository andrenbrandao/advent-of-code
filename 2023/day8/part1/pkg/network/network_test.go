package network

import (
	"testing"
)

func TestNode(t *testing.T) {

	t.Run("starts with null left", func(t *testing.T) {
		node := NewNode("AAA", nil, nil)

		got := node.Left()
		var want *Node

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("starts with null right", func(t *testing.T) {
		node := NewNode("AAA", nil, nil)

		got := node.Right()
		var want *Node

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("accepts left and right nodes", func(t *testing.T) {
		node := NewNode("AAA", NewNode("BBB", nil, nil), NewNode("CCC", nil, nil))

		left := node.Left()
		right := node.Right()

		if left == nil || right == nil {
			t.Errorf("left: %v, right: %v want them not nil", left, right)
		}
	})

	t.Run("compares two nodes", func(t *testing.T) {
		tests := []struct {
			node1 *Node
			node2 *Node
			want  bool
		}{
			{NewNode("AAA", nil, nil), NewNode("AAA", nil, nil), true},
			{NewNode("AAA", nil, nil), NewNode("BBB", nil, nil), false},
		}

		for _, tt := range tests {

			got := tt.node1.Equals(tt.node2)
			want := tt.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
}
