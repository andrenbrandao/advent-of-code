package camelcards

import (
	"reflect"
	"testing"
)

func TestGame(t *testing.T) {
	t.Run("sorts all the cards by ranks", func(t *testing.T) {
		tests := []struct {
			hands []string
			want  []string
		}{
			{
				[]string{"24443", "23432", "23456", "AAAAA", "AAAAQ", "23332"},
				[]string{"AAAAA", "AAAAQ", "33322", "44432", "43322", "65432"},
			},
		}

		for _, tt := range tests {
			game := NewGame(tt.hands)
			game.Sort()
			got := game
			want := NewGame(tt.want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
}
