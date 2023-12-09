package interval

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

// Joins two intervals from the beginning and end, including empty spaces
func (i *Interval) Join(other *Interval) *Interval {
	newStart := min(i.start, other.start)
	newEnd := max(i.end, other.end)
	return NewIntervalFromStartEnd(newStart, newEnd)
}

func (i *Interval) Minus(other *Interval) []*Interval {
	intersection := i.Intersection(other)

	if intersection == nil {
		return []*Interval{i}
	}

	// intersection is inside
	// split it into two
	if i.start < intersection.start && i.end > intersection.end {
		intervalA := NewIntervalFromStartEnd(i.start, intersection.start-1)
		intervalB := NewIntervalFromStartEnd(intersection.end+1, i.end)
		return []*Interval{intervalA, intervalB}
	}

	// removes right side
	if i.start < intersection.start {
		return []*Interval{NewIntervalFromStartEnd(i.start, intersection.start-1)}
	}

	// removes left side
	if i.end > intersection.end {
		return []*Interval{NewIntervalFromStartEnd(intersection.end+1, i.end)}
	}

	// equal
	// return empty array of intervals
	return []*Interval{}
}

func (i *Interval) Intersection(other *Interval) *Interval {
	startA := i.Start()
	endA := i.End()
	startB := other.Start()
	endB := other.End()

	startIntersect := max(startA, startB)
	endIntersect := min(endA, endB)

	// do not intersect
	if startIntersect > endIntersect {
		return nil
	}

	return NewIntervalFromStartEnd(startIntersect, endIntersect)
}
