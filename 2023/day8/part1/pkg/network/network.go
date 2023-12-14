package network

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
