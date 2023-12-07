package almanac

import (
	"fmt"
	"strconv"
	"strings"
)

type OptimizedMap struct {
	dataPoints  []string
	internalMap map[int]int
}

func NewOptimizedMap(lines []string) *OptimizedMap {
	aMap := OptimizedMap{dataPoints: lines, internalMap: make(map[int]int)}

	return &aMap
}

func (m *OptimizedMap) From(src int) int {
	fmt.Println("Creating map...", src)
	// m.calculateMap(src)

	val, ok := m.internalMap[src]

	if !ok {
		return src
	}

	return val
}

func (m *OptimizedMap) calculateMap(src int) {
	for _, line := range m.dataPoints {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		dataDst, _ := strconv.Atoi(fields[0])
		dataSrc, _ := strconv.Atoi(fields[1])
		aRange, _ := strconv.Atoi(fields[2])

		if src >= dataSrc && src <= dataSrc+aRange-1 {
			transformationDiff := dataDst - dataSrc
			m.internalMap[src] = src + transformationDiff
		}
	}
}
