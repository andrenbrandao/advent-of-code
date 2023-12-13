package camelcards

import (
	"reflect"
	"testing"
)

func TestGame(t *testing.T) {

	t.Run("sorts all the cards by ranks, from weakest to strongest", func(t *testing.T) {

		tests := []struct {
			hands []string
			want  []string
		}{
			{
				[]string{"24443", "23432", "23456", "AAAAA", "AAAAQ", "23332"},
				[]string{"23456", "23432", "24443", "23332", "AAAAQ", "AAAAA"},
			},
			{
				[]string{
					"AAAAA",
					"22222",
					"AAAAK",
					"22223",
					"AAAKK",
					"22233",
					"AAAKQ",
					"22234",
					"AAKKQ",
					"22334",
					"AAKQJ",
					"22345",
					"AKQJT",
					"23456",
				},
				[]string{
					"23456",
					"AKQJT",
					"22345",
					"AAKQJ",
					"22334",
					"AAKKQ",
					"22234",
					"AAAKQ",
					"22233",
					"AAAKK",
					"22223",
					"AAAAK",
					"22222",
					"AAAAA",
				},
			},
			// {
			// 	[]string{
			// 		"32T3K",
			// 		"T55J5",
			// 		"KK677",
			// 		"KTJJT",
			// 		"QQQJA",
			// 	},
			// 	[]string{
			// 		"32T3K",
			// 		"KK677",
			// 		"T55J5",
			// 		"QQQJA",
			// 		"KTJJT",
			// 	},
			// },
		}

		for _, tt := range tests {
			bids := []string{}
			for range tt.hands {
				bids = append(bids, "1")
			}

			game := NewGame(tt.hands, bids)
			game.Sort()
			got := game.hands
			want := NewGame(tt.want, bids).hands

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})

	t.Run("calculates the winnings by adding the multiplication of each bid with the rank", func(t *testing.T) {
		tests := []struct {
			input string
			want  int
		}{
			{
				`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`,
				6440,
			},
			{`AAAAA 2
22222 3
AAAAK 5
22223 7
AAAKK 11
22233 13
AAAKQ 17
22234 19
AAKKQ 23
22334 29
AAKQJ 31
22345 37
AKQJT 41
23456 43
`,
				1343,
			}}

		for _, tt := range tests {
			game := NewGameFromInput(tt.input)
			got := game.Winnings()
			want := tt.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
}
