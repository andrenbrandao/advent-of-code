package trebuchet

import (
	"strconv"
	"unicode"
)

type Trebuchet struct {
	calibrationDocument []string
}

func NewTrebuchet(calibrationDocument []string) *Trebuchet {
	return &Trebuchet{calibrationDocument}
}

func (t *Trebuchet) CalibrationValues() []int {
	result := make([]int, len(t.calibrationDocument))
	for i, s := range t.calibrationDocument {
		firstInt := -1
		lastInt := -1

		for _, c := range s {
			if !unicode.IsDigit(c) {
				continue
			}

			if firstInt == -1 {
				firstInt, _ = strconv.Atoi(string(c))
			}

			lastInt, _ = strconv.Atoi(string(c))
		}

		result[i] = firstInt*10 + lastInt
	}

	return result
}

func (t *Trebuchet) Sum() int {
	calibrationValues := t.CalibrationValues()
	sum := 0

	for _, val := range calibrationValues {
		sum += val
	}

	return sum
}
