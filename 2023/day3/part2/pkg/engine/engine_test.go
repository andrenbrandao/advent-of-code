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
		partNumbers := []int{467, 35, 633, 617, 592, 755, 664, 598}
		var want []*Part
		for _, partNumber := range partNumbers {
			p := Part(partNumber)
			want = append(want, &p)
		}

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

	t.Run("gets the gears in the schematics", func(t *testing.T) {
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
		got := len(engineSchematic.Gears())
		want := 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("sum of gear ratios", func(t *testing.T) {
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
		got := engineSchematic.SumGearRatios()
		want := 467835

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestGear(t *testing.T) {
	t.Run("gear ratio", func(t *testing.T) {
		parts := []*Part{NewPart(147), NewPart(285)}
		gear, _ := NewGear(parts)

		got := gear.GearRatio()
		want := 41895

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
