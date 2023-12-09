package almanac

import (
	"fmt"
	"sort"
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

func (i *Interval) Join(other *Interval) *Interval {
	newStart := min(i.start, other.start)
	newEnd := max(i.end, other.end)
	return NewIntervalFromStartEnd(newStart, newEnd)
}

func (i *Interval) Minus(other *Interval) *Interval {
	intersection := i.Intersection(other)

	// to the left
	if i.start < intersection.start {
		return NewIntervalFromStartEnd(i.start, intersection.start-1)
	}

	// to the right
	return NewIntervalFromStartEnd(intersection.end+1, i.end)
}

func (i *Interval) Intersection(other *Interval) *Interval {
	startA := i.Start()
	endA := i.End()
	startB := other.Start()
	endB := other.End()

	if startA > startB {
		startA, startB = startB, startA
		endA, endB = endB, endA
	}

	startIntersect := max(startA, startB)
	endIntersect := min(endA, endB)

	return NewIntervalFromStartEnd(startIntersect, endIntersect)
}

type IntervalMap struct {
	interval    *Interval
	internalMap *OptimizedMap
}

func NewIntervalMap(lines []string) *IntervalMap {
	intervals := []*Interval{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		dataSrc, _ := strconv.Atoi(fields[1])
		aRange, _ := strconv.Atoi(fields[2])

		interval := NewIntervalFromRange(dataSrc, aRange)
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
		interval:    joinedInterval,
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

func (im *IntervalMap) Transform(srcInterval *Interval) []*Interval {
	intervalA := srcInterval
	intervalB := im.interval

	startA := intervalA.Start()
	endA := intervalA.End()
	startB := intervalB.Start()
	endB := intervalB.End()

	fmt.Println("Interval A: ", startA, endA)
	fmt.Println("Interval B: ", startB, endB)

	// sort them
	if startA > startB {
		fmt.Println("swapped")
		intervalA, intervalB = intervalB, intervalA
		startA, startB = startB, startA
		endA, endB = endB, endA
	}

	startIntersect := max(startA, startB)
	endIntersect := min(endA, endB)

	fmt.Println("Intersect: ", startIntersect, endIntersect)

	// non-overlapping
	// no need to create others
	if startIntersect > endIntersect {
		fmt.Println("Non-overlapping")
		newStart := im.internalMap.From(srcInterval.Start())
		newEnd := im.internalMap.From(srcInterval.End())

		newInterval := NewIntervalFromStartEnd(newStart, newEnd)

		return []*Interval{newInterval}
	}

	startA = srcInterval.start
	endA = srcInterval.end

	// we mapped to inside the intersection
	// return the intersection mapped to the new values
	if startA >= startIntersect && endA <= endIntersect {
		newStart := im.internalMap.From(srcInterval.Start())
		newEnd := im.internalMap.From(srcInterval.End())

		newInterval := NewIntervalFromStartEnd(newStart, newEnd)

		return []*Interval{newInterval}
	}

	// has an intersection
	// split into intervals
	// B should not be part of the new intervals
	intersection := NewIntervalFromStartEnd(startIntersect, endIntersect)
	newA := srcInterval.Minus(intersection)
	startIntersect = im.internalMap.From(startIntersect)
	endIntersect = im.internalMap.From(endIntersect)
	mappedIntersection := NewIntervalFromStartEnd(startIntersect, endIntersect)

	return []*Interval{newA, mappedIntersection}
}

type IntervalAlmanac struct {
	seedIntervals []*Interval
	intervalMaps  []*IntervalMap
}

func NewIntervalAlmanac(input string) *IntervalAlmanac {
	lines := strings.Split(input, "\n")
	intervalMaps := make([]*IntervalMap, 0)
	seedIntervals := []*Interval{}

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

func extractSeedIntervals(line string) []*Interval {
	s := strings.Split(line, ":")
	fields := strings.Fields(s[1])
	intervals := []*Interval{}

	for i := 0; i < len(fields); i += 2 {
		start, _ := strconv.Atoi(fields[i])
		aRange, _ := strconv.Atoi(fields[i+1])
		newInterval := NewIntervalFromRange(start, aRange)
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

func (a *IntervalAlmanac) Locations() []*Interval {
	intervals := a.seedIntervals

	printIntervals(a.seedIntervals)

	for i := 0; i < 5; i++ {
		aMap := a.intervalMaps[i]

		fmt.Println()
		currentIntervals := []*Interval{}
		for _, interval := range intervals {
			for _, mappedInterval := range aMap.Transform(interval) {
				fmt.Println(mappedInterval)
				currentIntervals = append(currentIntervals, mappedInterval)
			}
		}

		printIntervals(currentIntervals)
		intervals = currentIntervals
	}

	return intervals
}

func (a *IntervalAlmanac) LowestLocation() int {
	locations := a.Locations()

	sort.Slice(locations, func(i, j int) bool {
		return locations[i].Start() < locations[j].Start()
	})

	// printIntervals(locations)
	return locations[0].Start()
}

func printIntervals(intervals []*Interval) {
	fmt.Printf("Current Intervals: ")

	for _, interval := range intervals {
		fmt.Printf("%d %d\n", interval.Start(), interval.End())
	}
}
