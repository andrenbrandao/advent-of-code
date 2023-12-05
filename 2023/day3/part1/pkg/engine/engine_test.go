package engine

import (
	"reflect"
	"testing"
)

func TestEngineSchematic(t *testing.T) {
	t.Run("extracts only parts next to symbols from the schematic", func(t *testing.T) {
		input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

		engineSchematic := NewEngineSchematic(input)
		got := engineSchematic.Parts()
		want := []Part{467, 35, 633, 617, 592, 755, 664, 598}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("gets sum of all parts", func(t *testing.T) {
		input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

		engineSchematic := NewEngineSchematic(input)
		got := engineSchematic.SumParts()
		want := 4361

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("ignores line breaks", func(t *testing.T) {
		input := `100
200`

		engineSchematic := NewEngineSchematic(input)
		got := engineSchematic.SumParts()
		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("works for one line", func(t *testing.T) {
		input := `503+`

		engineSchematic := NewEngineSchematic(input)
		got := engineSchematic.SumParts()
		want := 503

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
