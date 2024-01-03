package main

import (
	"day10/pkg/maze"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	maze := maze.NewMazeFromString(string(input))
	fmt.Println(maze.NumberEnclosedTiles())
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
