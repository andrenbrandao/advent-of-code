package engine

import (
	"strconv"
	"strings"
	"unicode"
)

type EngineSchematic struct {
	schematic string
}

type Part int

func NewEngineSchematic(input string) *EngineSchematic {
	return &EngineSchematic{input}
}

func NewPart(val int) *Part {
	part := Part(val)
	return &part
}

func (e *EngineSchematic) Parts() []Part {
	lines := strings.Split(e.schematic, "\n")
	var parts []Part

	for linePos, line := range lines {
		n := len(line)

		for i := 0; i < n; i++ {
			j := i

			if unicode.IsDigit(rune(line[j])) {
				for j < n && unicode.IsDigit(rune(line[j])) {
					j++
				}

				if !e.hasNeighborSymbol(i, j-1, linePos, lines) {
					i = j
					continue
				}

				partInt, _ := strconv.Atoi(line[i:j])
				parts = append(parts, *NewPart(partInt))

				i = j
			}
		}
	}

	return parts
}

func (e *EngineSchematic) hasNeighborSymbol(left, right, linePos int, lines []string) bool {
	isSymbol := func(i, j int, lines []string) bool {
		width := len(lines[0])
		height := len(lines)

		if i < 0 || i >= height || j < 0 || j >= width {
			return false
		}

		c := lines[i][j]
		return !unicode.IsDigit(rune(c)) && c != '.'
	}

	for j := left - 1; j <= right+1; j++ {
		for _, i := range []int{linePos - 1, linePos + 1} {
			if isSymbol(i, j, lines) {
				return true
			}
		}
	}

	return isSymbol(linePos, left-1, lines) || isSymbol(linePos, right+1, lines)
}
