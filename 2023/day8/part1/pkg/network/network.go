package network

import "strings"

type Node struct {
	id    string
	left  *Node
	right *Node
}

func NewNode(id string, node1 *Node, node2 *Node) *Node {
	return &Node{id: id, left: node1, right: node2}
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) Equals(other *Node) bool {
	return n.id == other.id
}

type Graph struct {
	steps string
	nodes []*Node
}

func NewGraphFromInstructions(input string) *Graph {
	lines := strings.Split(input, "\n")
	var steps string
	nodes := []*Node{}

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		if i == 0 {
			steps = line
			continue
		}

		s := strings.Split(line, "=")
		nodeId := strings.Trim(s[0], " ")
		node := NewNode(nodeId, nil, nil)
		nodes = append(nodes, node)
	}

	return &Graph{steps, nodes}
}

func (g *Graph) Count() int {
	return len(g.nodes)
}
