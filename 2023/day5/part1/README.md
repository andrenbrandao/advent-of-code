
# Algorithm

We want to find the paths from the seeds to the locations
going through maps. And we want to get the location with the lowest
number that was reached by a seed.

There are 2 ways:

- Start from a seed, find the locations and then sort them
- Start from the lowest location and go back the path trying to reach a seed

Depending on how many seeds we have, the second approach might be more efficient.

Let's take S as the number of seeds and L as the number of locations.

- If S << L, we should opt for Option 1
- If S >> L, we should opt for Option 2

Now, how do we get from the seeds to the locations?

This is a graph problem. We need to create a graph starting from the seeds and reaching
the locations.

## Steps

### Option 1

- For each seed, create each map pointing from and to the same number.
- Read each map data and update those maps with the source and destination
- For each seed, walk the graph and save the locations
- Sort the locations
- Return the lowest of them

I am going to assume that the maps data do not overlap, so you cannot
have a source going to multiple destinations.

#### Time and Space Complexities

- Iterating over each seed: O(S)
- Creation of a map: O(R) being R the range of possible source and destinations
- Walking the graph: O(S*M) being M the number of maps
- Sorting the locations: O(L*logL)

Therefore, the time complexity would be O(S+R+S*M+L*logL).

Space will also be O(R) if we create all the maps, and the sorting will take O(logL) depending
on the sorting algorithm.

#### Optimizations

There are some optimizations to be made:

1 - Skipping the creating of the maps for each value in the range.

This can take a lot of memory and is CPU intensive if the ranges are long.
A better approach would be to find the next step by calculating on the fly the
next destination, giving the current source.

2 - Large number of locations can slow down the algorithm

This can be optimized by following Option 2 approach and start going backwards
in the graph.

## OOP

Now, let's think of the objects and methods we need to create to handle this code.
We could design it by creating pure functions, but following OOP can make the code
easier to read and easier to change later.

### Almanac

Takes the puzzle input and creates the corresponding classes. Also returns the result.

### Seed, Soil, Fertilizer, Water, Light, Temperature, Humidity, Location

All of these are just integers, but will be created as types so that we protect our maps
from mapping incorrect objects.

### Maps

Each map will map a type to another. It will containg a `source`, `destination` and `range`.

### Graph

A graph will be constructed from Maps, Seeds and will return an array of Locations.
Then, the Almanac can return the lowest number of location.
