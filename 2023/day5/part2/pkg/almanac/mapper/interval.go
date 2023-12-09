package mapper

import (
	intvl "day5/pkg/almanac/interval"
	"sort"
	"strconv"
	"strings"
)

type IntervalMap struct {
	intervals   []*intvl.Interval
	internalMap *OptimizedMap
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

	aMap := NewOptimizedMap(lines)

	return &IntervalMap{
		intervals:   intervals,
		internalMap: aMap,
	}
}

// Transform create multiple intersections by doing srcInterval Minus the mapIntervals
// mapping those values to the new map and returning them
func (im *IntervalMap) Transform(srcInterval *intvl.Interval) []*intvl.Interval {
	intersections := []*intvl.Interval{}
	mappedIntersections := []*intvl.Interval{}

	for _, interval := range im.intervals {
		intersection := srcInterval.Intersection(interval)
		if intersection == nil {
			continue
		}
		intersections = append(intersections, intersection)

		newStart := im.internalMap.Transform(intersection.Start())
		newEnd := im.internalMap.Transform(intersection.End())

		mappedIntersections = append(mappedIntersections, intvl.NewIntervalFromStartEnd(newStart, newEnd))
	}

	notMappedIntervals := []*intvl.Interval{srcInterval}
	for _, i := range intersections {
		current := []*intvl.Interval{}

		for _, notMapped := range notMappedIntervals {
			current = append(current, notMapped.Minus(i)...)
		}

		notMappedIntervals = current
	}

	mappedIntervals := []*intvl.Interval{}
	mappedIntervals = append(mappedIntervals, mappedIntersections...)
	mappedIntervals = append(mappedIntervals, notMappedIntervals...)

	sort.Slice(mappedIntervals, func(i, j int) bool {
		return mappedIntervals[i].Start() < mappedIntervals[j].Start()
	})

	return mappedIntervals
}
