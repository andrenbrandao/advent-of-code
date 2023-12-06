package cards

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	id           int
	prizeNumbers []int
	numbers      []int
}

func NewCard(input string) *Card {
	cardNumber := extractCardNumber(input)
	prizeNumbers := extractPrizeNumbers(input)
	numbers := extractNumbers(input)

	return &Card{
		id:           cardNumber,
		prizeNumbers: prizeNumbers,
		numbers:      numbers,
	}
}

func extractCardNumber(input string) int {
	s := strings.Split(input, ":")
	cardStr := s[0]
	s = strings.Split(cardStr, " ")
	cardNumber, err := strconv.Atoi(string(s[1]))

	if err != nil {
		log.Fatal(err)
	}

	return cardNumber
}

func extractPrizeNumbers(input string) []int {
	s := strings.Split(input, ":")
	s = strings.Split(s[1], "|")
	prizeNumbersStr := strings.Fields(s[0])

	var prizeNumbers []int

	for _, n := range prizeNumbersStr {
		nInt, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		prizeNumbers = append(prizeNumbers, nInt)
	}

	return prizeNumbers
}

func extractNumbers(input string) []int {
	s := strings.Split(input, ":")
	s = strings.Split(s[1], "|")
	numbersStr := strings.Fields(s[1])

	var numbers []int

	for _, n := range numbersStr {
		nInt, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, nInt)
	}

	return numbers
}

func (c *Card) WinningNumbers() []int {
	var winningNumbers []int

	for _, prizeNumber := range c.prizeNumbers {
		for _, number := range c.numbers {
			if number == prizeNumber {
				winningNumbers = append(winningNumbers, number)
			}
		}
	}

	return winningNumbers
}

func (c *Card) Points() int {
	winningNumbers := c.WinningNumbers()
	n := float64(len(winningNumbers))

	if n == 0 {
		return 0
	}

	return int(math.Pow(2, n-1))
}
