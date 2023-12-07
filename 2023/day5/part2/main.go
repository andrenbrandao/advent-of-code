package main

import (
	"day5/pkg/almanac"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	almanac := almanac.NewAlmanac(string(input), &almanac.RangeSeedExtractor{})
	fmt.Println(almanac.LowestLocation())
}
