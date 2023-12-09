package almanac

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGraphAlmanac(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

	t.Run("returns the location mapped from a seed", func(t *testing.T) {
		almanac := NewGraphAlmanac(input, &DefaultSeedExtractor{})
		got := almanac.Locations()
		want := []int{82, 43, 86, 35}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("returns the lowest location", func(t *testing.T) {
		almanac := NewGraphAlmanac(input, &DefaultSeedExtractor{})
		got := almanac.LowestLocation()
		want := 35

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("returns the lowest location for a Range Seed Extractor", func(t *testing.T) {
		almanac := NewGraphAlmanac(input, &RangeSeedExtractor{})
		got := almanac.LowestLocation()
		want := 46

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns the lowest location using the reverse path", func(t *testing.T) {
		almanac := NewGraphAlmanac(input, &RangeSeedExtractor{})
		got := almanac.OptimizedLowestLocation()
		want := 46

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

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
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(35, 40), NewIntervalFromStartEnd(35, 37)},
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(38, 40), nil},
			{NewIntervalFromStartEnd(10, 37), NewIntervalFromStartEnd(0, 12), NewIntervalFromStartEnd(10, 12)},
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
			interval := NewIntervalFromStartEnd(tt.sourceIntervalStartEnd[0], tt.sourceIntervalStartEnd[1])
			got := intervalMap.Transform(interval)
			for _, g := range got {
				fmt.Println("XXXXX")
				fmt.Println(g.start, g.end)
				fmt.Println("XXXXX")
			}

			var want []*Interval
			for _, intervalRange := range tt.destStartEnd {
				destInterval := NewIntervalFromStartEnd(intervalRange[0], intervalRange[1])
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
			interval := NewIntervalFromStartEnd(tt.sourceIntervalStartEnd[0], tt.sourceIntervalStartEnd[1])
			got := intervalMap.Transform(interval)

			var want []*Interval
			for _, intervalRange := range tt.destStartEnd {
				destInterval := NewIntervalFromStartEnd(intervalRange[0], intervalRange[1])
				want = append(want, destInterval)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		}

	})
}

func TestIntervalAlmanac(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

	t.Run("returns the location with the lowest value", func(t *testing.T) {
		almanac := NewIntervalAlmanac(input)
		got := almanac.LowestLocation()
		want := 46

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
