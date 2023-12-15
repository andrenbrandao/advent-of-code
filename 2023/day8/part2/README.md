# The Problem

We now have to start on multiple nodes that end in 'A' and can only stop when all the nodes stop at a node ending in 'Z' at the same time.

## Algorithm

- Go through all nodes, if they end in 'A', add as a source
- Add these sources to a queue

While we don´t have all in ending in 'Z':

- Pop all sources from the queue
- Execute the same instruction for all of them
- Add their next nodes to the queue
