package main

import (
	"day6/pkg/race"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	raceRecords := race.NewRaceRecords(string(input))
	fmt.Println(raceRecords.MultipliedNumberWays())
}
