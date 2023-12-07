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

type Almanac struct {
	seeds []int
	maps  []*Map
}

func NewAlmanac(input string) *Almanac {
	lines := strings.Split(input, "\n")
	seeds := []int{}
	maps := []*Map{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seeds = extractSeeds(line)
		}

		if strings.Contains(line, "map:") {
			var aMap *Map
			i, aMap = extractMap(i+1, lines)
			maps = append(maps, aMap)
		}
	}

	return &Almanac{
		seeds: seeds,
		maps:  maps,
	}
}

func extractSeeds(line string) []int {
	s := strings.Split(line, ":")
	fields := strings.Fields(s[1])
	seeds := []int{}

	for _, field := range fields {
		fieldInt, _ := strconv.Atoi(field)
		seeds = append(seeds, fieldInt)
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
