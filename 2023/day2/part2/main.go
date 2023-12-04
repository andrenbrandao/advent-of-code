package main

import (
	"bufio"
	"day2/pkg/games"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	var gamesInput []string
	for reader.Scan() {
		gamesInput = append(gamesInput, reader.Text())
	}

	gameRecords := games.NewGameRecords(gamesInput)
	fmt.Println(gameRecords.SumPowerSetCubes())
}
