package almanac

import (
	intvl "day5/pkg/almanac/interval"
	"day5/pkg/almanac/mapper"
	"sort"
	"strconv"
	"strings"
)

type IntervalMap struct {
	joinedInterval *intvl.Interval
	intervals      []*intvl.Interval
	internalMap    *mapper.OptimizedMap
}

func NewIntervalMap(lines []string) *IntervalMap {
	intervals := []*intvl.Interval{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		dataSrc, _ := strconv.Atoi(fields[1])
		aRange, _ := strconv.Atoi(fields[2])

		interval := intvl.NewIntervalFromRange(dataSrc, aRange)
		intervals = append(intervals, interval)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start() < intervals[j].Start()
	})

	joinedInterval := intervals[0]

	for i := 1; i < len(intervals); i++ {
		joinedInterval = joinedInterval.Join(intervals[i])
	}

	aMap := mapper.NewOptimizedMap(lines)

	return &IntervalMap{
		joinedInterval: joinedInterval,
		intervals:      intervals,
		internalMap:    aMap,
	}
}

/*

case 1
|-----|
        |------|


case 2
          |-----|
|-----|


so, first we need to sort it to eliminate multiple cases
** Sorting won´t solve everything **
** It will actually confuse what is the source and the filter **
** Sorting makes it easier to get the intersection **

case 1 -> creates 1 interval
|-----|
        |------|

case 2 -> creates 3 intervals
|-----|
    |------|

case 3
|------| -> creates 3 intervals
  |--|

case 4
|-----| -> creates 1 interval
|-----|

case 5
|-----| -> creates 2 intervals
  |---|

case 6 -> returns the same interval
      |----|
|---|

case 7 -> return the intersection and the right side

  |----|
|---|


startIntersect = max(startA, startB)
endIntersect = min(endA, endB)

if startIntersect > endIntersect: {
	NOT VALID, just return the mapping to the same values
}

OTHERWISE:

newStartA = startA
newEndA = min(endA, startIntersect-1)

newStartB = max(startB, endIntersect+1)
newEndB = max(endB, endA)


case 3:
	returns the newA, intersection and newB

case 4:
	if intersection is the same as A, return A

case 7:

*/

func (im *IntervalMap) Transform(srcInterval *intvl.Interval) []*intvl.Interval {
	intervalA := srcInterval
	intervalB := im.joinedInterval

	startA := intervalA.Start()
	endA := intervalA.End()
	startB := intervalB.Start()
	endB := intervalB.End()

	// sort them
	if startA > startB {
		intervalA, intervalB = intervalB, intervalA
		startA, startB = startB, startA
		endA, endB = endB, endA
	}

	startIntersect := max(startA, startB)
	endIntersect := min(endA, endB)

	// non-overlapping
	// no need to create others
	if startIntersect > endIntersect {
		newStart := im.internalMap.From(srcInterval.Start())
		newEnd := im.internalMap.From(srcInterval.End())

		newInterval := intvl.NewIntervalFromStartEnd(newStart, newEnd)

		return []*intvl.Interval{newInterval}
	}

	// create multiple intersections by doing Minus the intervals
	// map those values to the new map
	// return them
	mappedIntervals := []*intvl.Interval{}
	for _, interval := range im.intervals {
		intersection := srcInterval.Intersection(interval)
		if intersection != nil {
			newStart := im.internalMap.From(intersection.Start())

			newEnd := im.internalMap.From(intersection.End())
			mappedIntervals = append(mappedIntervals, intvl.NewIntervalFromStartEnd(newStart, newEnd))
		}
	}

	mappedIntervals = append(mappedIntervals, srcInterval.Minus(im.joinedInterval)...)

	sort.Slice(mappedIntervals, func(i, j int) bool {
		return mappedIntervals[i].Start() < mappedIntervals[j].Start()
	})

	return mappedIntervals
}

type IntervalAlmanac struct {
	seedIntervals []*intvl.Interval
	intervalMaps  []*IntervalMap
}

func NewIntervalAlmanac(input string) *IntervalAlmanac {
	lines := strings.Split(input, "\n")
	intervalMaps := make([]*IntervalMap, 0)
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
			var aMap *IntervalMap
			i, aMap = extractIntervalMap(i+1, lines)
			intervalMaps = append(intervalMaps, aMap)
		}
	}

	return &IntervalAlmanac{
		seedIntervals: seedIntervals,
		intervalMaps:  intervalMaps,
	}
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

func extractIntervalMap(i int, lines []string) (int, *IntervalMap) {
	top := i
	bottom := i

	for len(lines[bottom]) > 0 {
		bottom++
	}

	aMap := NewIntervalMap(lines[top:bottom])

	return bottom, aMap
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
