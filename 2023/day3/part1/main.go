package main

import (
	"bufio"
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

	var input []string
	for reader.Scan() {
		input = append(input, reader.Text())
	}

	fmt.Println(input)
}
