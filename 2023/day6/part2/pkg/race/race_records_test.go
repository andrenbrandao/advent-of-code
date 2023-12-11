package race

import (
	"reflect"
	"testing"
)

func TestRaceRecords(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200
` // this now means Time: 71530 and Distance: 940200

	t.Run("returns number of ways to win all races", func(t *testing.T) {
		want := []int{71503}
		raceRecords := NewRaceRecords(input)
		got := raceRecords.WaysToWinAllRaces()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("multiplies the number of ways to win", func(t *testing.T) {
		want := 71503
		raceRecords := NewRaceRecords(input)
		got := raceRecords.MultipliedNumberWays()

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
