package boat

import (
	"day6/pkg/race/types"
	"testing"
)

func TestBoat(t *testing.T) {

	t.Run("runs as certaing distance given a charge", func(t *testing.T) {
		distanceTests := []struct {
			charge    types.Charge
			totalTime types.Time
			want      types.Distance
		}{
			{charge: 0, totalTime: 7, want: 0},
			{charge: 1, totalTime: 7, want: 6},
			{charge: 3, totalTime: 7, want: 12},
		}

		for _, tt := range distanceTests {
			boat := NewBoat(tt.charge)
			got := boat.Distance(tt.totalTime - types.Time(tt.charge))

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		}
	})
}
