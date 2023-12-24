package main

import (
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

	fmt.Println(input)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
