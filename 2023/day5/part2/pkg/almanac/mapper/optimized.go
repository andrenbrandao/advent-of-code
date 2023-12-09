package mapper

import (
	"strconv"
	"strings"
)

type OptimizedMap struct {
	dataPoints         []string
	internalMap        map[int]int
	internalReverseMap map[int]int
}

func NewOptimizedMap(lines []string) *OptimizedMap {
	aMap := OptimizedMap{
		dataPoints:         lines,
		internalMap:        make(map[int]int),
		internalReverseMap: make(map[int]int)}

	return &aMap
}

func (m *OptimizedMap) From(src int) int {
	m.calculateMap(src)

	val, ok := m.internalMap[src]

	if !ok {
		return src
	}

	return val
}

func (m *OptimizedMap) Keys() []int {
	keys := make([]int, 0, len(m.internalMap))
	for k := range m.internalMap {
		keys = append(keys, k)
	}

	return keys
}

func (m *OptimizedMap) ReverseKeys() []int {
	keys := make([]int, 0, len(m.internalReverseMap))
	for k := range m.internalReverseMap {
		keys = append(keys, k)
	}

	return keys
}

func (m *OptimizedMap) FromReverse(src int) int {
	m.calculateReverseMap(src)

	val, ok := m.internalReverseMap[src]

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

func (m *OptimizedMap) calculateReverseMap(src int) {
	for _, line := range m.dataPoints {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		dataSrc, _ := strconv.Atoi(fields[0])
		dataDst, _ := strconv.Atoi(fields[1])
		aRange, _ := strconv.Atoi(fields[2])

		if src >= dataSrc && src <= dataSrc+aRange-1 {
			transformationDiff := dataDst - dataSrc
			m.internalReverseMap[src] = src + transformationDiff
		}
	}
}
