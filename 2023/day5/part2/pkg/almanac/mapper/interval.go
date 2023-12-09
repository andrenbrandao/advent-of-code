package mapper

import (
	intvl "day5/pkg/almanac/interval"
	"sort"
	"strconv"
	"strings"
)

type IntervalMap struct {
	joinedInterval *intvl.Interval
	intervals      []*intvl.Interval
	internalMap    *OptimizedMap
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

	aMap := NewOptimizedMap(lines)

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
	// create multiple intersections by doing Minus the intervals
	// map those values to the new map
	// return them
	mappedIntervals := []*intvl.Interval{}
	for _, interval := range im.intervals {
		intersection := srcInterval.Intersection(interval)
		if intersection != nil {
			newStart := im.internalMap.Transform(intersection.Start())

			newEnd := im.internalMap.Transform(intersection.End())
			mappedIntervals = append(mappedIntervals, intvl.NewIntervalFromStartEnd(newStart, newEnd))
		}
	}

	mappedIntervals = append(mappedIntervals, srcInterval.Minus(im.joinedInterval)...)

	sort.Slice(mappedIntervals, func(i, j int) bool {
		return mappedIntervals[i].Start() < mappedIntervals[j].Start()
	})

	return mappedIntervals
}
