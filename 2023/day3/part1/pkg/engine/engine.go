package engine

import (
	"strconv"
	"strings"
	"unicode"
)

type EngineSchematic struct {
	schematic string
}

func NewEngineSchematic(input string) *EngineSchematic {
	return &EngineSchematic{input}
}

func (e *EngineSchematic) Parts() []int {
	lines := strings.Split(e.schematic, "\n")
	var parts []int

	for _, line := range lines {
		n := len(line)

		for i := 0; i < n; i++ {
			j := i

			if unicode.IsDigit(rune(line[j])) {
				for j < n && unicode.IsDigit(rune(line[j])) {
					j++
				}

				partInt, _ := strconv.Atoi(line[i:j])
				parts = append(parts, partInt)

				i = j
			}
		}
	}

	return parts
}
