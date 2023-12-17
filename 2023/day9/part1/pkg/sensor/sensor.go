package sensor

import (
	"strconv"
	"strings"
)

type Sensor struct {
	values []int
}

func NewSensor(input string) *Sensor {
	values := []int{}

	vStrings := strings.Fields(input)
	for _, val := range vStrings {
		vInt, _ := strconv.Atoi(val)
		values = append(values, vInt)
	}

	return &Sensor{values}
}

func (s *Sensor) Next() int {
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

	return recHelper(s.values)
}
