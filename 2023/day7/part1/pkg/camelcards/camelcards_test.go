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

	t.Run("compares two hands", func(t *testing.T) {
		tests := []struct {
			hand1 string
			hand2 string
			want  bool // winner first
		}{
			{"AAAAA", "AAAAQ", true},  // five of kind
			{"AAAAA", "QQQQQ", true},  // five of kind
			{"QQQQQ", "AAAAA", false}, // five of kind
			{"AAAAQ", "23332", true},  // four of kind
			{"23332", "13332", true},  // full house
			{"13332", "23432", true},  // three of kind
			{"23432", "23456", true},  // two pair
			{"23456", "33456", true},  // high card
			{"23456", "23457", false}, // high card
		}

		for _, tt := range tests {
			hand1 := NewHand(tt.hand1)
			hand2 := NewHand(tt.hand2)
			got := hand1.StrongerThan(hand2)

			if got != tt.want {
				t.Errorf("%s, %s: got %v, want %v", tt.hand1, tt.hand2, got, tt.want)
			}
		}
	})
}
