package almanac

import (
	"sort"
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

type OptimizedMap struct {
	dataPoints  []string
	internalMap map[int]int
}

func NewOptimizedMap(lines []string) *OptimizedMap {
	aMap := OptimizedMap{dataPoints: lines, internalMap: make(map[int]int)}

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

type Almanac struct {
	seeds []int
	maps  []*OptimizedMap
}

type SeedExtractor interface {
	extract(line string) []int
}

func NewAlmanac(input string, seedExtractor SeedExtractor) *Almanac {
	lines := strings.Split(input, "\n")
	seeds := []int{}
	maps := []*OptimizedMap{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seeds = seedExtractor.extract(line)
		}

		if strings.Contains(line, "map:") {
			var aMap *OptimizedMap
			i, aMap = extractOptimizedMap(i+1, lines)
			maps = append(maps, aMap)
		}
	}

	return &Almanac{
		seeds: seeds,
		maps:  maps,
	}
}

type DefaultSeedExtractor struct{}

func (e *DefaultSeedExtractor) extract(line string) []int {
	s := strings.Split(line, ":")
	fields := strings.Fields(s[1])
	seeds := []int{}

	for _, field := range fields {
		fieldInt, _ := strconv.Atoi(field)
		seeds = append(seeds, fieldInt)
	}

	return seeds
}

type RangeSeedExtractor struct{}

func (e *RangeSeedExtractor) extract(line string) []int {
	s := strings.Split(line, ":")
	fields := strings.Fields(s[1])
	seeds := []int{}

	for i := 0; i < len(fields); i += 2 {
		seed, _ := strconv.Atoi(fields[i])
		aRange, _ := strconv.Atoi(fields[i+1])

		for k := seed; k < seed+aRange-1; k++ {
			seeds = append(seeds, k)
		}
	}

	return seeds
}

func extractMap(i int, lines []string) (int, *Map) {
	top := i
	bottom := i
	for len(lines[bottom]) > 0 {
		bottom++
	}

	m := NewMap(lines[top:bottom])
	return bottom, m
}

func extractOptimizedMap(i int, lines []string) (int, *OptimizedMap) {
	top := i
	bottom := i
	for len(lines[bottom]) > 0 {
		bottom++
	}

	m := NewOptimizedMap(lines[top:bottom])
	return bottom, m
}

func (a *Almanac) Locations() []int {
	locations := []int{}

	for _, seed := range a.seeds {
		dest := seed

		for _, aMap := range a.maps {
			dest = aMap.From(dest)
		}

		locations = append(locations, dest)
	}

	return locations
}

func (a *Almanac) LowestLocation() int {
	locations := a.Locations()
	sort.Slice(locations, func(i, j int) bool {
		return locations[i] < locations[j]
	})

	return locations[0]
}
