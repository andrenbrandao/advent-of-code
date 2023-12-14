package main

import (
	"day8/pkg/network"
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

	graph := network.NewGraphFromInstructions(string(input))
	fmt.Println(graph.StepsCount())
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
