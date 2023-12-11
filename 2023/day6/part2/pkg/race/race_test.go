package race

import (
	"day6/pkg/race/types"
	"testing"
)

func TestRace(t *testing.T) {

	t.Run("gives ways to win for given time and record distances", func(t *testing.T) {
		raceTests := []struct {
			maxTime        types.Time
			recordDistance types.Distance
			want           int
		}{
			{maxTime: 7, recordDistance: 9, want: 4},
		}

		for _, tt := range raceTests {
			race := NewRace(tt.maxTime, tt.recordDistance)
			got := race.WaysToWin()

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		}
	})
}
