package engine

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type EngineSchematic struct {
	schematic string
	parts     []*Part
	gears     []*Gear
}

type Node interface {
	Neighbors() []*Node
}

type Part int

type Root struct{}

func (n *Root) Neighbors() []*Node {
	return []*Node{}
}

func NewEngineSchematic(input string) *EngineSchematic {
	return &EngineSchematic{input, []*Part{}, []*Gear{}}
}

func NewPart(val int) *Part {
	part := Part(val)
	return &part
}

func (p *Part) ToInt() int {
	return int(*p)
}

func (e *EngineSchematic) Parts() []*Part {
	// return cached parts
	if len(e.parts) > 0 {
		return e.parts
	}

	parts := e.extractParts()

	// cache parts
	e.parts = parts
	return parts
}

func (e *EngineSchematic) extractParts() []*Part {
	lines := strings.Split(e.schematic, "\n")
	var parts []*Part

	for linePos, line := range lines {
		n := len(line)

		for i := 0; i < n; i++ {
			if !unicode.IsDigit(rune(line[i])) {
				continue
			}

			j := i
			for j < n && unicode.IsDigit(rune(line[j])) {
				j++
			}

			if e.hasNeighborSymbol(i, j-1, linePos, lines) {
				partInt, _ := strconv.Atoi(line[i:j])
				parts = append(parts, NewPart(partInt))
			}

			i = j
		}
	}
	return parts
}

func (e *EngineSchematic) hasNeighborSymbol(left, right, linePos int, lines []string) bool {
	width := len(lines[0])
	height := len(lines)

	isSymbol := func(i, j int, lines []string) bool {
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

func (e *EngineSchematic) SumParts() int {
	sum := 0
	for _, p := range e.Parts() {
		sum += p.ToInt()
	}

	return sum
}

func (e *EngineSchematic) Gears() []*Gear {
	// return cached parts
	if len(e.gears) > 0 {
		return e.gears
	}

	gears := e.extractGears()

	// cache gears
	e.gears = gears
	return gears
}

type Gear struct {
	parts []*Part
}

func NewGear(parts []*Part) (*Gear, error) {
	if len(parts) != 2 {
		return nil, errors.New("Invalid number of parts")
	}
	return &Gear{parts}, nil
}

func (e *EngineSchematic) extractGears() []*Gear {
	lines := strings.Split(e.schematic, "\n")
	var gears []*Gear

	for linePos, line := range lines {
		n := len(line)

		for j := 0; j < n; j++ {
			if line[j] != '*' {
				continue
			}

			parts := e.neighborParts(j, linePos, lines)
			gear, err := NewGear(parts)

			if err == nil {
				gears = append(gears, gear)
			}
		}
	}

	return gears
}

func (e *EngineSchematic) neighborParts(pos, linePos int, lines []string) []*Part {
	width := len(lines[0])
	height := len(lines)
	visited := make(map[[2]int]bool)

	isVisited := func(i, j int) bool {
		return visited[[2]int{i, j}]
	}

	getPart := func(i, j int, lines []string) *Part {
		if i < 0 || i >= height || j < 0 || j >= width || !unicode.IsDigit(rune(lines[i][j])) {
			return nil
		}

		for j >= 0 && unicode.IsDigit(rune(lines[i][j])) {
			j--
		}

		left := j + 1
		right := left

		for right < width && unicode.IsDigit(rune(lines[i][right])) {
			visited[[2]int{i, right}] = true
			right++
		}

		valInt, _ := strconv.Atoi(string(lines[i][left:right]))
		return NewPart(valInt)
	}

	parts := []*Part{}

	for j := pos - 1; j <= pos+1; j++ {
		for i := linePos - 1; i <= linePos+1; i++ {
			if i == linePos && j == pos {
				continue
			}

			if !isVisited(i, j) {
				part := getPart(i, j, lines)
				if part != nil {
					parts = append(parts, part)
				}
			}
		}
	}

	return parts
}
