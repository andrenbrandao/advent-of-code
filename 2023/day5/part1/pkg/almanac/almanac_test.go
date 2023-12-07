package almanac

import "testing"

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

		myMap := NewMap(input)

		for _, tt := range mapTests {
			if myMap[tt.source] != tt.want {
				t.Errorf("got %v want %v", tt.source, tt.want)
			}
		}
	})
}
