package main

import (
	"day7/pkg/camelcards"
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

	game := camelcards.NewGameFromInput(string(input))
	fmt.Println(game.Winnings())
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
