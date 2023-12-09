package almanac

import (
	intvl "day5/pkg/almanac/interval"
	"day5/pkg/almanac/mapper"
	"sort"
	"strconv"
	"strings"
)

type IntervalAlmanac struct {
	seedIntervals []*intvl.Interval
	intervalMaps  []*mapper.IntervalMap
}

func NewIntervalAlmanac(input string) *IntervalAlmanac {
	lines := strings.Split(input, "\n")
	intervalMaps := make([]*mapper.IntervalMap, 0)
	seedIntervals := []*intvl.Interval{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seedIntervals = extractSeedIntervals(line)
		}

		if strings.Contains(line, "map:") {
			var aMap *mapper.IntervalMap
			i, aMap = extractIntervalMap(i+1, lines)
			intervalMaps = append(intervalMaps, aMap)
		}
	}

	return &IntervalAlmanac{
		seedIntervals: seedIntervals,
		intervalMaps:  intervalMaps,
	}
}

func (a *IntervalAlmanac) Locations() []*intvl.Interval {
	intervals := a.seedIntervals

	for i := 0; i < len(a.intervalMaps); i++ {
		aMap := a.intervalMaps[i]

		currentIntervals := []*intvl.Interval{}
		for _, interval := range intervals {
			for _, mappedInterval := range aMap.Transform(interval) {
				currentIntervals = append(currentIntervals, mappedInterval)
			}
		}

		intervals = currentIntervals
	}

	return intervals
}

func (a *IntervalAlmanac) LowestLocation() int {
	locations := a.Locations()

	sort.Slice(locations, func(i, j int) bool {
		return locations[i].Start() < locations[j].Start()
	})

	return locations[0].Start()
}

func extractSeedIntervals(line string) []*intvl.Interval {
	s := strings.Split(line, ":")
	fields := strings.Fields(s[1])
	intervals := []*intvl.Interval{}

	for i := 0; i < len(fields); i += 2 {
		start, _ := strconv.Atoi(fields[i])
		aRange, _ := strconv.Atoi(fields[i+1])
		newInterval := intvl.NewIntervalFromRange(start, aRange)
		intervals = append(intervals, newInterval)
	}

	return intervals
}

func extractIntervalMap(i int, lines []string) (int, *mapper.IntervalMap) {
	top := i
	bottom := i

	for len(lines[bottom]) > 0 {
		bottom++
	}

	aMap := mapper.NewIntervalMap(lines[top:bottom])

	return bottom, aMap
}
