# The Problem

We have a graph/map with a starting node AAA and a destination node ZZZ. There is also a set of instructions "LLR" that dictates if we should walk to the left or right along the graph.

How many steps are necessary to reach the destination node ZZZ?

## Algorithm

- Read the steps to be followed
- Create a graph from the nodes given
- Starting from the node AAA, repeat the steps until we reach node ZZZ
- Count these number of steps and return

## Use Cases

### Node

- Has a left and right node

### Graph

- Contains multiple nodes
- Should be able to find a node in O(1), like in a hashmap
- Iterate over the map following the steps to reach the destination

## Design

```
Node
constructor:
id string
left *Node
right *Node

methods:
+ Equals(*Node) bool
+ Left() *Node
+ Right() *Node


---

Graph:
constructor:
[]*Node

methods:
+ StepsCount(sequence string) int


```

## Todo

### Node

- [x] Create Node
- [x] Left
- [x] Right
- [x] Equals

### Graph

- [ ] Create Graph struct
- [ ] Map of nodes to other nodes
- [ ] StepsCount
