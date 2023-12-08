package almanac

import (
	"fmt"
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func NewInterval(input string) *Interval {
	fields := strings.Fields(input)
	start, _ := strconv.Atoi(fields[0])
	aRange, _ := strconv.Atoi(fields[1])

	return NewIntervalFromRange(start, aRange)
}

func NewIntervalFromRange(start int, aRange int) *Interval {
	end := start + aRange - 1

	return &Interval{
		start: start,
		end:   end,
	}
}

func NewIntervalFromStartEnd(start int, end int) *Interval {
	return &Interval{
		start: start,
		end:   end,
	}
}

func (i *Interval) Start() int {
	return i.start
}

func (i *Interval) End() int {
	return i.end
}

type IntervalMap struct {
	interval    *Interval
	internalMap *OptimizedMap
}

func NewIntervalMap(input string) *IntervalMap {
	fields := strings.Fields(input)
	dataSrc, _ := strconv.Atoi(fields[1])
	aRange, _ := strconv.Atoi(fields[2])

	interval := NewIntervalFromRange(dataSrc, aRange)
	aMap := NewOptimizedMap([]string{input})

	return &IntervalMap{
		interval:    interval,
		internalMap: aMap,
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
*/

func (im *IntervalMap) Transform(srcInterval *Interval) []*Interval {
	intervalA := srcInterval
	intervalB := im.interval

	startA := intervalA.Start()
	endA := intervalA.End()
	startB := intervalB.Start()
	endB := intervalB.End()

	// sort them
	if startA > startB {
		intervalA, intervalB = intervalB, intervalA
	}

	startIntersect := max(startA, startB)
	endIntersect := min(endA, endB)

	// non-overlapping
	// no need to create others
	if startIntersect > endIntersect {
		newStart := im.internalMap.From(srcInterval.Start())
		newEnd := im.internalMap.From(srcInterval.End())

		newInterval := NewIntervalFromStartEnd(newStart, newEnd)

		return []*Interval{newInterval}
	}

	// the intersection is the same, map only the original interval
	if startA == startIntersect && endA == endIntersect {
		newStart := im.internalMap.From(srcInterval.Start())
		newEnd := im.internalMap.From(srcInterval.End())

		newInterval := NewIntervalFromStartEnd(newStart, newEnd)

		return []*Interval{newInterval}
	}

	// has an intersection
	// split into intervals
	newStartA := startA
	newEndA := min(endA, startIntersect-1)

	newStartB := max(startB, endIntersect+1)
	newEndB := max(endB, endA)

	newA := NewIntervalFromStartEnd(newStartA, newEndA)
	newB := NewIntervalFromStartEnd(newStartB, newEndB)

	startIntersect = im.internalMap.From(startIntersect)
	endIntersect = im.internalMap.From(endIntersect)
	intersection := NewIntervalFromStartEnd(startIntersect, endIntersect)

	return []*Interval{newA, intersection, newB}
}

func NewSeedIntervalMap(input string) *IntervalMap {
	fields := strings.Fields(input)
	dataSrc, _ := strconv.Atoi(fields[0])
	aRange, _ := strconv.Atoi(fields[1])
	dataDst := dataSrc
	mapInput := fmt.Sprintf("%d %d %d", dataDst, dataSrc, aRange)

	interval := NewIntervalFromRange(dataSrc, aRange)
	aMap := NewOptimizedMap([]string{mapInput})

	return &IntervalMap{
		interval:    interval,
		internalMap: aMap,
	}
}

type IntervalAlmanac struct {
	seedsInteval []*Interval
}

// func NewIntervalAlmanac(input string) *IntervalAlmanac {
// 	lines := strings.Split(input, "\n")
// 	var seedsInterval []*Interval

// 	for i := 0; i < len(lines); i++ {
// 		line := lines[i]

// 		if len(line) == 0 {
// 			continue
// 		}

// 		if strings.HasPrefix(line, "seeds:") {
// 			seedsInterval = extractSeedIntervals(line)
// 		}
// 	}

// 	return &IntervalAlmanac{
// 		seedsInput: seedsInput,
// 		maps:       maps,
// 	}
// }
