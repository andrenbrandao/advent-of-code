package almanac

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

	fmt.Println("Finished extracting seeds")
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
