package sensor

import (
	"testing"
)

func TestSensor(t *testing.T) {

	t.Run("predicts the next value", func(t *testing.T) {
		tests := []struct {
			input string
			want  int
		}{
			{"0 3 6 9 12 15", 18},
			{"1 3 6 10 15 21", 28},
			{"10 13 16 21 30 45", 68},
			{"-6 -4 -2 0 2 4", 6},
		}
		for _, tt := range tests {
			sensor := NewSensor(tt.input)

			got := sensor.Next(0)
			want := tt.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})

	t.Run("sums all the next values", func(t *testing.T) {
		tests := []struct {
			input string
			want  int
		}{
			{`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
			`, 114},
		}

		for _, tt := range tests {
			sensor := NewSensor(tt.input)

			got := sensor.Sum()
			want := tt.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
}
