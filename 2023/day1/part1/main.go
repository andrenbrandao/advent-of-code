package main

import (
	"bufio"
	"day1/pkg/trebuchet"
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

	var document []string
	for reader.Scan() {
		document = append(document, reader.Text())
	}

	trebuchet := trebuchet.NewTrebuchet(document)
	fmt.Println(trebuchet.Sum())
}
