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
	// part2_suboptimal(input)
	part2(input)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part1(input []byte) {
	almanac := almanac.NewGraphAlmanac(string(input), &almanac.DefaultSeedExtractor{})
	fmt.Println(almanac.LowestLocation())
}

func part2_suboptimal(input []byte) {
	almanac := almanac.NewGraphAlmanac(string(input), &almanac.RangeSeedExtractor{})
	fmt.Println(almanac.OptimizedLowestLocation())
}

func part2(input []byte) {
	almanac := almanac.NewIntervalAlmanac(string(input))
	fmt.Println(almanac.LowestLocation())
}
