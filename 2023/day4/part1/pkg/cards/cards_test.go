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
}
