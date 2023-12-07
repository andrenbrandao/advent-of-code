package almanac

import (
	"strconv"
	"strings"
)

type Map map[int]int

func NewMap(input string) Map {
	lines := strings.Split(input, "\n")
	aMap := Map{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		destination, _ := strconv.Atoi(fields[0])
		source, _ := strconv.Atoi(fields[1])
		aRange, _ := strconv.Atoi(fields[2])

		for i := 0; i < aRange; i++ {
			aMap[source+i] = destination + i
		}
	}

	return aMap
}
