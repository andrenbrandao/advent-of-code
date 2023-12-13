package camelcards

import "testing"

func TestCard(t *testing.T) {
	t.Run("compares card with another", func(t *testing.T) {
		cardTests := []struct {
			cardType1 rune
			cardType2 rune
			want      bool
		}{
			{'A', 'K', true},
			{'Q', '3', true},
			{'T', '2', true},
			{'2', 'A', false},
			{'9', 'T', false},
			{'2', '9', false},
		}

		for _, tt := range cardTests {
			card1 := NewCard(CardType(tt.cardType1))
			card2 := NewCard(CardType(tt.cardType2))
			got := card1.StrongerThan(card2)

			if got != tt.want {
				t.Errorf("Card1: %c, Card2: %c, got %v, want %v", tt.cardType1, tt.cardType2, got, tt.want)
			}
		}
	})
}
