package main

import (
	"day4/pkg/cards"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	scratchCards := cards.NewScratchCards(string(input))
	fmt.Println(scratchCards.TotalScratchCardsDP())
}
