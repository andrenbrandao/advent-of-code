package almanac

import (
	"strconv"
	"strings"
)

type Map struct {
	internalMap map[int]int
}

func NewMap(lines []string) *Map {
	aMap := Map{internalMap: make(map[int]int)}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		destination, _ := strconv.Atoi(fields[0])
		source, _ := strconv.Atoi(fields[1])
		aRange, _ := strconv.Atoi(fields[2])

		for i := 0; i < aRange; i++ {
			aMap.internalMap[source+i] = destination + i
		}
	}

	return &aMap
}

func (m *Map) From(src int) int {
	val, ok := m.internalMap[src]

	if !ok {
		return src
	}

	return val
}
