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

func TestGraph(t *testing.T) {
	input :=
		`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`
	t.Run("has count of nodes", func(t *testing.T) {
		graph := NewGraphFromInstructions(input)

		got := graph.Count()
		want := 7

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("count number of steps to find destination", func(t *testing.T) {
		graph := NewGraphFromInstructions(input)

		got := graph.StepsCount()
		want := 2

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("find number of steps from multiple sources", func(t *testing.T) {
		input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`
		graph := NewGraphFromInstructions(input)

		got := graph.StepsCount()
		want := 6

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
