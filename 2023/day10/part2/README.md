# The Problem

We want to find the number of tiles that are INSIDE the loop.

## Algorithm

There are different ways of solving it.

- One it using [Pick's Theorem](https://en.wikipedia.org/wiki/Pick%27s_theorem) to calculate the area of the polygon.
- The [Scanline Polygon Fill Algorithm](https://www.educative.io/answers/what-is-scanline-fill-algorithm#:~:text=The%20scanline%20fill%20algorithm%20works,minimum%20to%20maximum%20y%2Dvalues.)
- The [Nonzero-Rule](https://en.wikipedia.org/wiki/Nonzero-rule) algorithm.
- The [Even-odd Rule](https://en.wikipedia.org/wiki/Even%E2%80%93odd_rule)

I will use the latter.

```
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........

```

- We start by using a similar algorithm to the FarthestSteps to iterate over the Path and create a new atrix with the tiles that are present as "#".
- Then, we iterate over it again and every time we find a tile that is not "#", we go either left or right and count the number of tiles it intersects.
- If we count an odd number of tiles, it is inside. Otherwise, it is outside.
- Change those to either "I" or "O" so we can print it in the end. Keep the count of each.

Note: we can only consider tiles that are vertical for the intersect. Horizontal ones can't be used. However, how do we define which one the "S" will be?

## Results

After implementing the Even-odd rule, one of the test cases failed. That happened because it can have even number of tiles to the side and be inside.

```
...........
FS.........
||F-7......
|||.|......
|||.|......
|LJxL7F7...
|....LJ|...
|......|...
L------J...

```

That happens because at some specific turns it ends up going to the right instead of crossing the path. If it crossed, it would have to come back, always counting 2 tiles.

L--7 -> 1
L--J -> 2
F--7 -> 2
F--J -> 1
F- -> 1
