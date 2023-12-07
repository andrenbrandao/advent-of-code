package almanac

import (
	"strconv"
	"strings"
)

type SeedExtractor interface {
	extract(line string) []int
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
