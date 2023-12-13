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

func TestHand(t *testing.T) {
	t.Run("defines the type of the hand", func(t *testing.T) {
		tests := []struct {
			name string
			hand string
			want int
		}{
			{"Five of a kind", "AAAAA", FIVE_OF_KIND},
			{"Four of a kind", "AAAAQ", FOUR_OF_KIND},
			{"Full house", "23332", FULL_HOUSE},
			{"Three of a kind", "13332", THREE_OF_KIND},
			{"Two pair", "23432", TWO_PAIR},
			{"High card", "23456", HIGH_CARD},
		}

		for _, tt := range tests {
			got := NewHand(tt.hand).Type()

			if got != tt.want {
				t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
			}
		}
	})
}
