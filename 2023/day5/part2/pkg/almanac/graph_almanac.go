package almanac

import (
	"day5/pkg/almanac/mapper"
	"day5/pkg/almanac/seeds"
	"sort"
	"strings"
)

type GraphAlmanac struct {
	seedsInput    string
	maps          []*mapper.OptimizedMap
	seedExtractor seeds.SeedExtractor
}

func NewGraphAlmanac(input string, seedExtractor seeds.SeedExtractor) *GraphAlmanac {
	lines := strings.Split(input, "\n")
	var seedsInput string
	maps := []*mapper.OptimizedMap{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seedsInput = line
		}

		if strings.Contains(line, "map:") {
			var aMap *mapper.OptimizedMap
			i, aMap = extractOptimizedMap(i+1, lines)
			maps = append(maps, aMap)
		}
	}

	return &GraphAlmanac{
		seedsInput:    seedsInput,
		maps:          maps,
		seedExtractor: seedExtractor,
	}
}

func (a *GraphAlmanac) Seeds() []int {
	return a.seedExtractor.Extract(a.seedsInput)
}

func (a *GraphAlmanac) Locations() []int {
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

func (a *GraphAlmanac) LowestLocation() int {
	locations := a.Locations()
	sort.Slice(locations, func(i, j int) bool {
		return locations[i] < locations[j]
	})

	return locations[0]
}

// Really slow
// Tries to traverse the graph in reverse from the lowest location
// but the numbers are really big, so this may take hours/days
func (a *GraphAlmanac) OptimizedLowestLocation() int {
	reversedMaps := a.maps
	for i, j := 0, len(reversedMaps)-1; i < j; i, j = i+1, j-1 {
		reversedMaps[i], reversedMaps[j] = reversedMaps[j], reversedMaps[i]
	}
	seedSet := seeds.NewSeedSet(a.seedsInput)

	for location := 0; location < 3530465412; location++ {
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
