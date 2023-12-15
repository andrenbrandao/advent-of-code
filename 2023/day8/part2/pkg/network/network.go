package network

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

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

type NodeStep struct {
	node *Node
	pos  int
}

func (ns *NodeStep) String() string {
	return fmt.Sprintf("%s, %d", ns.node, ns.pos)
}

func (g *Graph) StepsCount() int {
	srcNodes := g.sourceNodes()

	nodeSteps := []*NodeStep{}
	for _, node := range srcNodes {
		nodeSteps = append(nodeSteps, &NodeStep{node, 0})
	}

	numberOfSteps := 0

	outsideQueue := []*NodeStep{}
	outsideQueue = append(outsideQueue, nodeSteps...)
	insideQueue := []*NodeStep{}

	for {
		fmt.Println(outsideQueue)
		insideQueue = append(insideQueue, outsideQueue...)
		outsideQueue = outsideQueue[:0]

		count := 0
		for _, nodeStep := range insideQueue {
			if nodeStep.node.IsDestination() {
				count++
			}
		}

		if count == len(insideQueue) {
			return numberOfSteps
		}

		numberOfSteps++

		for len(insideQueue) > 0 {
			nodeStep := insideQueue[0]
			node := nodeStep.node
			stepPos := nodeStep.pos
			fmt.Println(nodeStep)

			insideQueue = insideQueue[1:]

			switch g.steps[stepPos] {
			case 'L':
				node = node.Left()
			case 'R':
				node = node.Right()
			default:
				log.Fatal("Wrong direction")
			}

			stepPos = (stepPos + 1) % len(g.steps)
			outsideQueue = append(outsideQueue, &NodeStep{node, stepPos})
		}
	}
}

func (g *Graph) sourceNodes() []*Node {
	nodes := []*Node{}

	for _, node := range g.nodes {
		if node.IsSource() {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
