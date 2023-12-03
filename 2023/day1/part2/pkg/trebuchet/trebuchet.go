package trebuchet

type Trebuchet struct {
	calibrationDocument []string
}

func NewTrebuchet(calibrationDocument []string) *Trebuchet {
	return &Trebuchet{calibrationDocument}
}

func (t *Trebuchet) CalibrationValues() []int {
	result := make([]int, len(t.calibrationDocument))

	for i, s := range t.calibrationDocument {
		value := t.extractValue(s)

		result[i] = value
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

func (*Trebuchet) extractValue(s string) int {
	tokenMap := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	firstInt := -1
	lastInt := -1
	n := len(s)

	// Time Complexity: O(n^2*t*k)
	// where n is the size of the word
	// t is the size of the longest token
	// k is the number of tokens
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			c := s[i : j+1]

			val, ok := tokenMap[string(c)]
			if !ok {
				continue
			}

			i = j + 1

			if firstInt == -1 {
				firstInt = val
			}

			lastInt = val
		}
	}

	return firstInt*10 + lastInt
}
