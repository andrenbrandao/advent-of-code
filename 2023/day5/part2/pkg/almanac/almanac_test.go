package almanac

import (
	"reflect"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("creates a map from an input", func(t *testing.T) {
		input := `0 15 37
37 52 2
39 0 15
`

		mapTests := []struct {
			source int
			want   int
		}{
			{15, 0},
			{16, 1},
			{51, 36},
			{52, 37},
			{0, 39},
		}

		myMap := NewMap(strings.Split(input, "\n"))

		for _, tt := range mapTests {
			got := myMap.From(tt.source)
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})

	t.Run("maps numbers to themselves", func(t *testing.T) {
		input := `0 15 37`

		mapTests := []struct {
			source int
			want   int
		}{
			{0, 0},
			{14, 14},
			{52, 52},
		}

		myMap := NewMap(strings.Split(input, "\n"))

		for _, tt := range mapTests {
			got := myMap.From(tt.source)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})
}

func TestOptimizedMap(t *testing.T) {
	t.Run("creates a map from an input", func(t *testing.T) {
		input := `0 15 37
37 52 2
39 0 15
`

		mapTests := []struct {
			source int
			want   int
		}{
			{15, 0},
			{16, 1},
			{51, 36},
			{52, 37},
			{0, 39},
		}

		myMap := NewOptimizedMap(strings.Split(input, "\n"))

		for _, tt := range mapTests {
			got := myMap.From(tt.source)
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})

	t.Run("maps numbers to themselves", func(t *testing.T) {
		input := `0 15 37`

		mapTests := []struct {
			source int
			want   int
		}{
			{0, 0},
			{14, 14},
			{52, 52},
		}

		myMap := NewOptimizedMap(strings.Split(input, "\n"))

		for _, tt := range mapTests {
			got := myMap.From(tt.source)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})
}

func TestAlmanac(t *testing.T) {
	t.Run("returns the location mapped from a seed", func(t *testing.T) {
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

		almanac := NewAlmanac(input, &DefaultSeedExtractor{})
		got := almanac.Locations()
		want := []int{82, 43, 86, 35}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("returns the lowest location", func(t *testing.T) {
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

		almanac := NewAlmanac(input, &DefaultSeedExtractor{})
		got := almanac.LowestLocation()
		want := 35

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("returns the lowest location for a Range Seed Extractor", func(t *testing.T) {
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

		almanac := NewAlmanac(input, &RangeSeedExtractor{})
		got := almanac.LowestLocation()
		want := 46

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns the lowest location using the reverse path", func(t *testing.T) {
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

		almanac := NewAlmanac(input, &RangeSeedExtractor{})
		got := almanac.OptimizedLowestLocation()
		want := 46

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
