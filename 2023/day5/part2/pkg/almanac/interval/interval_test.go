package interval

import (
	"reflect"
	"testing"
)

func TestInterval(t *testing.T) {
	t.Run("create an interval from a given start and range string", func(t *testing.T) {
		input := "10 37"

		interval := NewInterval(input)
		got := []int{interval.Start(), interval.End()}
		want := []int{10, 46}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("joins with another interval", func(t *testing.T) {
		interval1 := NewIntervalFromStartEnd(10, 37)
		interval2 := NewIntervalFromStartEnd(38, 50)
		got := interval1.Join(interval2)
		want := NewIntervalFromStartEnd(10, 50)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("subtract two intervals", func(t *testing.T) {
		intervalTests := []struct {
			interval1 *Interval
			interval2 *Interval
			want      []*Interval
		}{
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(35, 40), []*Interval{NewIntervalFromStartEnd(10, 34)}},
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(38, 40), []*Interval{NewIntervalFromStartEnd(10, 37)}},
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(0, 12), []*Interval{NewIntervalFromStartEnd(13, 37)}},
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(20, 30), []*Interval{NewIntervalFromStartEnd(10, 19), NewIntervalFromStartEnd(31, 37)}},
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(11, 36), []*Interval{NewIntervalFromStartEnd(10, 10), NewIntervalFromStartEnd(37, 37)}},
		}

		for _, tt := range intervalTests {
			got := tt.interval1.Minus(tt.interval2)
			want := tt.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		}
	})

	t.Run("intersects two intervals", func(t *testing.T) {
		intervalTests := []struct {
			interval1 *Interval
			interval2 *Interval
			want      *Interval
		}{
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(35, 40), NewIntervalFromStartEnd(35, 37)}, // overlap
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(0, 12), NewIntervalFromStartEnd(10, 12)},  // overlap
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(37, 38), NewIntervalFromStartEnd(37, 37)}, // overlap at a point
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(10, 37)}, // equal
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(38, 40), nil},                             // do not overlap
		}

		for _, tt := range intervalTests {
			got := tt.interval1.Intersection(tt.interval2)
			want := tt.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		}
	})
}
