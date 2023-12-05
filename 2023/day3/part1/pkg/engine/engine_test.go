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

}
