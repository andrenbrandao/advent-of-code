package main

import (
	"day3/pkg/engine"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	engineSchematic := engine.NewEngineSchematic(string(input))
	fmt.Println(engineSchematic.SumGearRatios())
}
