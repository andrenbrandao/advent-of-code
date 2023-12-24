# The Problem

We have a matrix with pipes and we want to find the longest distance from the S character. Need to create a graph that is a cycle.

## Algorithm

F----7 -> not a loop we are interested because there is no S
L----J

7    J
|    |
S----7
|    L--7
L-------J

### Assumptions

- There is only 1 cycle connecting to the character S. Otherwise, we would have more than one cycle and it goes against the problem's description.
- There can be other cycles but they won't be connected to S.
- The farthest point from "S" is always the same distance walking from both sides, which means is in the middle. We can get that by counting the total steps of the cycle and dividing by 2.
- Since the cycle is made by walking on vertical and horizontal axis, the number of steps is always even, so we can confidently divide them by 2.

### Considerations

- Can't use BFS initially because we need to go back to the valid path if we wind up in a character after S that does not complete the cycle. We need to use DFS.
- I would run DFS to create the Graph and then BFS to find the longest distance.

### How to

- Iterate over the matrix
- When we find "S", run a DFS to create the graph
  - Look only to the sides that the letter allow you to (S looks to any of them, 7 is left and down, etc)
  - If we cannot go to any of the sides, backtrack
  - Otherwise, mark the node as visited and keep searching
  - When we reach back to S, we have found the cycle

We keep track of the number of steps during the DFS. The total / 2 is the answer.

## Design

```
MazeFactory
- input string
+ create() *Maze

Maze
- aMap [][]Tile
+ StartingPos() Pos

Tile (interface)
+ Neighbors() []Pos

Ground
- pos Pos
+ Neighbors() []Pos

Pipe
- pos Pos
+ Neighbors() []Pos
```
