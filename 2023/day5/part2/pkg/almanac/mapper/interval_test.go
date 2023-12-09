package mapper

import (
	intvl "day5/pkg/almanac/interval"
	"reflect"
	"strings"
	"testing"
)

func TestIntervalMap(t *testing.T) {
	t.Run("intersects a given interval with the current map and transform into the next", func(t *testing.T) {
		intervalTests := []struct {
			mapInput               string
			sourceIntervalStartEnd []int
			destStartEnd           [][]int
		}{
			// dest, src, range
			{"0 15 37", []int{15, 51}, [][]int{[]int{0, 36}}},                               // same range
			{"0 15 37", []int{21, 41}, [][]int{[]int{6, 26}}},                               // inside intersection
			{"0 15 37", []int{14, 52}, [][]int{[]int{0, 36}, []int{14, 14}, []int{52, 52}}}, // outside intersection
		}

		for _, tt := range intervalTests {
			intervalMap := NewIntervalMap(strings.Split(tt.mapInput, "\n"))
			interval := intvl.NewIntervalFromStartEnd(tt.sourceIntervalStartEnd[0], tt.sourceIntervalStartEnd[1])
			got := intervalMap.Transform(interval)

			var want []*intvl.Interval
			for _, intervalRange := range tt.destStartEnd {
				destInterval := intvl.NewIntervalFromStartEnd(intervalRange[0], intervalRange[1])
				want = append(want, destInterval)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		}

	})

	t.Run("tries to match with the multiple intervals inside a map", func(t *testing.T) {
		intervalTests := []struct {
			mapInput               string
			sourceIntervalStartEnd []int
			destStartEnd           [][]int
		}{
			// dest, src, range
			{`0 15 37
			37 52 2`, []int{15, 53}, [][]int{[]int{0, 36}, []int{37, 38}}},
		}

		for _, tt := range intervalTests {
			intervalMap := NewIntervalMap(strings.Split(tt.mapInput, "\n"))
			interval := intvl.NewIntervalFromStartEnd(tt.sourceIntervalStartEnd[0], tt.sourceIntervalStartEnd[1])
			got := intervalMap.Transform(interval)

			var want []*intvl.Interval
			for _, intervalRange := range tt.destStartEnd {
				destInterval := intvl.NewIntervalFromStartEnd(intervalRange[0], intervalRange[1])
				want = append(want, destInterval)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		}

	})
}
