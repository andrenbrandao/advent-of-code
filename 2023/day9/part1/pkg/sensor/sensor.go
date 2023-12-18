package sensor

import (
	"strconv"
	"strings"
)

type Sensor struct {
	history [][]int
}

func NewSensor(input string) *Sensor {
	history := [][]int{}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		values := []int{}
		vStrings := strings.Fields(line)
		for _, val := range vStrings {
			vInt, _ := strconv.Atoi(val)
			values = append(values, vInt)
		}
		history = append(history, values)
	}

	return &Sensor{history}
}

func (s *Sensor) Next(day int) int {
	var recHelper func(arr []int) int
	recHelper = func(arr []int) int {
		zeroCount := 0
		for _, v := range arr {
			if v == 0 {
				zeroCount++
			}
		}

		if zeroCount == len(arr) {
			return 0
		}

		nextArr := []int{}
		for i := 1; i < len(arr); i++ {
			diff := arr[i] - arr[i-1]
			nextArr = append(nextArr, diff)
		}

		lastVal := arr[len(arr)-1]
		return lastVal + recHelper(nextArr)
	}

	return recHelper(s.history[day])
}

func (s *Sensor) Sum() int {
	total := 0

	for day := range s.history {
		val := s.Next(day)
		total += val
	}

	return total
}
