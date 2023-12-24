package maze

import (
	"reflect"
	"testing"
)

func TestPipe(t *testing.T) {

	t.Run("returns the next positions", func(t *testing.T) {
		tests := []struct {
			pipeType rune
			want     []Pos
		}{
			{'|', []Pos{{0, -1}, {0, 1}}},
			{'-', []Pos{{-1, 0}, {1, 0}}},
			{'L', []Pos{{0, -1}, {1, 0}}},
			{'J', []Pos{{-1, 0}, {0, -1}}},
			{'7', []Pos{{-1, 0}, {0, 1}}},
			{'F', []Pos{{1, 0}, {0, 1}}},
			{'S', []Pos{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}},
		}

		for _, tt := range tests {
			pipe := NewPipe(tt.pipeType, Pos{0, 0})

			got := pipe.Neighbors()
			want := tt.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("type %c, got %v, want %v", tt.pipeType, got, want)
			}
		}
	})

}