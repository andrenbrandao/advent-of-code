package camelcards

import (
	"reflect"
	"testing"
)

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
			{"23332", "24443", true},  // full house
			{"24443", "23432", true},  // three of kind
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

	t.Run("sorts a hand", func(t *testing.T) {
		tests := []struct {
			hand string
			want string
		}{
			{"QAAAA", "AAAAQ"},
			{"23332", "33322"},
			{"24443", "44432"},
		}

		for _, tt := range tests {
			hand := NewHand(tt.hand)
			hand.Sort()
			got := hand
			want := NewHand(tt.want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
}
