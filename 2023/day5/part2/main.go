package main

import (
	"day5/pkg/almanac"
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

	// part1(input)
	part2(input)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part1(input []byte) {
	almanac := almanac.NewAlmanac(string(input), &almanac.DefaultSeedExtractor{})
	fmt.Println(almanac.LowestLocation())
}

func part2(input []byte) {
	almanac := almanac.NewAlmanac(string(input), &almanac.RangeSeedExtractor{})
	fmt.Println(almanac.OptimizedLowestLocation())
}
