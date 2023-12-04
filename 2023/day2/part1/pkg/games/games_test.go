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
