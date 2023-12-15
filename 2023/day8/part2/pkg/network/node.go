package network

import "fmt"

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

func (n *Node) AddLeftChild(other *Node) {
	n.left = other
}

func (n *Node) AddRightChild(other *Node) {
	n.right = other
}

func (n *Node) IsSource() bool {
	return n.id[len(n.id)-1] == 'A'
}

func (n *Node) IsDestination() bool {
	return n.id[len(n.id)-1] == 'Z'
}

func (n *Node) String() string {
	return fmt.Sprintf("%s", n.id)
}
