package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
}
