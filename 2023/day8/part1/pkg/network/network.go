package network

import (
	"log"
	"regexp"
	"strings"
)

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

type Graph struct {
	steps       string
	nodes       []*Node
	idToNodeMap map[string]*Node
}

func NewGraphFromInstructions(input string) *Graph {
	lines := strings.Split(input, "\n")
	var steps string
	nodes := []*Node{}
	idToNodeMap := map[string]*Node{}

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		if i == 0 {
			steps = line
			continue
		}

		r, _ := regexp.Compile("\\w+")
		matches := r.FindAllString(line, -1)

		nodeId1 := matches[0]
		leftNodeId := matches[1]
		rightNodeId := matches[2]
		var node *Node
		var leftNode *Node
		var rightNode *Node

		if n, ok := idToNodeMap[nodeId1]; !ok {
			node = NewNode(nodeId1, nil, nil)
			idToNodeMap[nodeId1] = node
		} else {
			node = n
		}

		if n, ok := idToNodeMap[leftNodeId]; !ok {
			leftNode = NewNode(leftNodeId, nil, nil)
			idToNodeMap[leftNodeId] = leftNode
		} else {
			leftNode = n
		}

		if n, ok := idToNodeMap[rightNodeId]; !ok {
			rightNode = NewNode(rightNodeId, nil, nil)
			idToNodeMap[rightNodeId] = rightNode
		} else {
			rightNode = n
		}

		node.AddLeftChild(leftNode)
		node.AddRightChild(rightNode)
		nodes = append(nodes, node)
	}

	return &Graph{steps, nodes, idToNodeMap}
}

func (g *Graph) Count() int {
	return len(g.nodes)
}

func (g *Graph) StepsCount() int {
	source := "AAA"
	destination := "ZZZ"

	srcNode := g.idToNodeMap[source]
	dstNode := g.idToNodeMap[destination]

	count := 0
	node := srcNode
	for {
		for _, c := range g.steps {
			if node.Equals(dstNode) {
				return count
			}

			count++

			switch c {
			case 'L':
				node = node.Left()
				continue
			case 'R':
				node = node.Right()
				continue
			default:
				log.Fatal("Wrong direction")
			}
		}
	}
}
