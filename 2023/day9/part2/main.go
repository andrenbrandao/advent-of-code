package main

import (
	"day8/pkg/sensor"
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

	s := sensor.NewSensor(string(input))
	fmt.Println(s.SumAllPrev())
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
