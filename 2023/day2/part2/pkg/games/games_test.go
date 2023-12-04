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

	t.Run("returns a play with the minimum number of cubes to enable the two plays", func(t *testing.T) {
		validationTests := []struct {
			play1 string
			play2 string
			want  *Play
		}{
			{"3 blue, 4 red", "4 blue, 3 red", &Play{blue: 4, red: 4}},
			{"1 red, 2 green, 6 blue", "3 red, 3 blue", &Play{red: 3, green: 2, blue: 6}},
		}

		for _, tt := range validationTests {
			got := NewPlay(tt.play1).MinPlay(NewPlay(tt.play2))

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})
}

func TestGame(t *testing.T) {
	t.Run("creates game with an id", func(t *testing.T) {
		input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		game := NewGame(input)

		got := game.Id()
		want := 1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("creates game with multiple plays", func(t *testing.T) {
		input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		game := NewGame(input)

		got := len(game.Plays())
		want := 3

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("validates a game based on a valid play", func(t *testing.T) {
		validationTests := []struct {
			input string
			want  bool
		}{
			{"Game 1: 3 blue, 4 red", true},
			{"Game 2: 1 red, 2 green, 6 blue", true},
			{"Game 3: 12 red, 13 green, 14 blue", true},
			{"Game 4: 13 red, 13 green, 14 blue", false},
			{"Game 5: 12 red, 14 green, 14 blue", false},
			{"Game 6: 12 red, 13 green, 15 blue", false},
			{"Game 7: 1 red, 2 green, 6 blue; 12 red, 13 green, 15 blue", false},
		}

		validPlay := NewPlay("12 red, 13 green, 14 blue")

		for _, tt := range validationTests {
			got := NewGame(tt.input).Valid(validPlay)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})

	t.Run("returns the minimum valid play", func(t *testing.T) {
		validationTests := []struct {
			input string
			want  *Play
		}{
			{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", &Play{red: 4, blue: 6, green: 2}},
			{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", &Play{red: 1, blue: 4, green: 3}},
			{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", &Play{red: 20, blue: 6, green: 13}},
			{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", &Play{red: 14, blue: 15, green: 3}},
			{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", &Play{red: 6, blue: 2, green: 3}},
		}

		for _, tt := range validationTests {
			got := NewGame(tt.input).MinPlay()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})
}

func TestGameRecords(t *testing.T) {
	t.Run("valid games", func(t *testing.T) {
		validationTests := []struct {
			input []string
			want  []int
		}{
			{[]string{"Game 1: 3 blue, 4 red",
				"Game 2: 1 red, 2 green, 6 blue",
				"Game 3: 12 red, 13 green, 14 blue",
				"Game 4: 13 red, 13 green, 14 blue",
				"Game 5: 12 red, 14 green, 14 blue",
				"Game 6: 12 red, 13 green, 15 blue",
				"Game 7: 1 red, 2 green, 6 blue; 12 red, 13 green, 15 blue"},
				[]int{1, 2, 3}},
		}

		validPlay := NewPlay("12 red, 13 green, 14 blue")

		for _, tt := range validationTests {
			games := NewGameRecords(tt.input).ValidGames(validPlay)

			var got []int
			for _, game := range games {
				got = append(got, game.Id())
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})

	t.Run("sum of valid game ids", func(t *testing.T) {
		validationTests := []struct {
			input []string
			want  int
		}{
			{[]string{"Game 1: 3 blue, 4 red",
				"Game 2: 1 red, 2 green, 6 blue",
				"Game 3: 12 red, 13 green, 14 blue",
				"Game 4: 13 red, 13 green, 14 blue",
				"Game 5: 12 red, 14 green, 14 blue",
				"Game 6: 12 red, 13 green, 15 blue",
				"Game 7: 1 red, 2 green, 6 blue; 12 red, 13 green, 15 blue"},
				6},
		}

		validPlay := NewPlay("12 red, 13 green, 14 blue")

		for _, tt := range validationTests {
			got := NewGameRecords(tt.input).SumValidGameIds(validPlay)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		}
	})
}
