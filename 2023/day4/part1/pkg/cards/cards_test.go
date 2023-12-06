package cards

import (
	"reflect"
	"testing"
)

func TestCard(t *testing.T) {
	t.Run("winning numbers", func(t *testing.T) {
		input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"

		card := NewCard(input)
		got := card.WinningNumbers()
		want := []int{48, 83, 86, 17}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("calculates points", func(t *testing.T) {
		pointTests := []struct {
			input string
			want  int
		}{
			{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
			{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
			{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
			{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
			{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
			{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
		}

		for _, tt := range pointTests {
			card := NewCard(tt.input)
			got := card.Points()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})
}

func TestScratchCards(t *testing.T) {
	t.Run("total points", func(t *testing.T) {
		input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

		scratchCards := NewScratchCards(input)
		got := scratchCards.TotalPoints()
		want := 13

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
