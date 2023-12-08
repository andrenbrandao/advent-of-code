package almanac

import (
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

func (im *IntervalMap) Transform(interval *Interval) *Interval {
	newStart := im.internalMap.From(interval.Start())
	newEnd := im.internalMap.From(interval.End())

	return NewIntervalFromStartEnd(newStart, newEnd)
}

// type IntervalAlmanac struct {
// 	seedsInteval []*Interval
// }

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
