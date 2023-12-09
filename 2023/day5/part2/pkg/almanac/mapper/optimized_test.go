package mapper

import (
	"strings"
	"testing"
)

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
