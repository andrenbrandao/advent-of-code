package almanac

import (
	"fmt"
	"sort"
	"strings"
)

type Almanac struct {
	seedsInput    string
	maps          []*OptimizedMap
	seedExtractor SeedExtractor
}

func NewAlmanac(input string, seedExtractor SeedExtractor) *Almanac {
	lines := strings.Split(input, "\n")
	var seedRanges string
	maps := []*OptimizedMap{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seedRanges = line
		}

		if strings.Contains(line, "map:") {
			var aMap *OptimizedMap
			i, aMap = extractOptimizedMap(i+1, lines)
			maps = append(maps, aMap)
		}
	}

	return &Almanac{
		seedsInput:    seedRanges,
		maps:          maps,
		seedExtractor: seedExtractor,
	}
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

func (a *Almanac) Seeds() []int {
	return a.seedExtractor.extract(a.seedsInput)
}

func (a *Almanac) Locations() []int {
	locations := []int{}

	for _, seed := range a.Seeds() {
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

func (a *Almanac) OptimizedLowestLocation() int {
	reversedMaps := a.maps
	for i, j := 0, len(reversedMaps)-1; i < j; i, j = i+1, j-1 {
		reversedMaps[i], reversedMaps[j] = reversedMaps[j], reversedMaps[i]
	}
	seedSet := NewSeedSet(a.seedsInput)

	for location := 0; location < 3530465412; location++ {
		fmt.Println(location)
		dest := location
		for _, aMap := range reversedMaps {
			dest = aMap.FromReverse(dest)
		}
		if seedSet.Has(dest) {
			return location
		}
	}

	return -1
}
