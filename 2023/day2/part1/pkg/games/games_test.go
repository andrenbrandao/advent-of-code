package games

import (
	"reflect"
	"testing"
)

func TestPlay(t *testing.T) {
	t.Run("parses a play", func(t *testing.T) {
		input := "3 blue, 4 red"

		play := NewPlay(input)
		got := []int{play.Blue(), play.Red()}
		want := []int{3, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("validates it against a valid play", func(t *testing.T) {
		validationTests := []struct {
			input string
			want  bool
		}{
			{"3 blue, 4 red", true},
			{"1 red, 2 green, 6 blue", true},
			{"12 red, 13 green, 14 blue", true},
			{"13 red, 13 green, 14 blue", false},
			{"12 red, 14 green, 14 blue", false},
			{"12 red, 13 green, 15 blue", false},
		}

		validPlay := NewPlay("12 red, 13 green, 14 blue")

		for _, tt := range validationTests {
			got := NewPlay(tt.input).Valid(validPlay)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})
}

func TestGame(t *testing.T) {
	t.Run("creates game with multiple plays", func(t *testing.T) {
		input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		game := NewGame(input)

		got := len(game.Plays())
		want := 3

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
