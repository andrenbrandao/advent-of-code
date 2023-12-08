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
	end := start + aRange - 1

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
