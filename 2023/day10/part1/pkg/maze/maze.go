package maze

import (
	"errors"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Pipe struct {
	pipeType rune
	pos      Pos
}

func NewPipe(pipeType rune, pos Pos) *Pipe {
	return &Pipe{pipeType, pos}
}

func (p *Pipe) Neighbors() []Pos {
	top := Pos{p.pos.x, p.pos.y - 1}
	bottom := Pos{p.pos.x, p.pos.y + 1}
	left := Pos{p.pos.x - 1, p.pos.y}
	right := Pos{p.pos.x + 1, p.pos.y}

	switch p.pipeType {
	case '|':
		return []Pos{top, bottom}
	case '-':
		return []Pos{left, right}
	case 'L':
		return []Pos{top, right}
	case 'J':
		return []Pos{left, top}
	case '7':
		return []Pos{left, bottom}
	case 'F':
		return []Pos{right, bottom}
	case 'S':
		return []Pos{left, top, right, bottom}
	}

	return []Pos{}
}

func (p *Pipe) Char() rune {
	return p.pipeType
}

type Ground struct {
	pos Pos
}

func NewGround(pos Pos) *Ground {
	return &Ground{pos}
}

func (g *Ground) Neighbors() []Pos {
	return []Pos{}
}

func (g *Ground) Char() rune {
	return '.'
}

type Tile interface {
	Neighbors() []Pos
	Char() rune
}
type Maze struct {
	tiles [][]Tile
}

func NewMazeFromString(input string) *Maze {
	tiles := [][]Tile{}
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		currentTiles := []Tile{}

		for j, c := range line {
			if c == '.' {
				currentTiles = append(currentTiles, NewGround(Pos{i, j}))
				continue
			}

			currentTiles = append(currentTiles, NewPipe(c, Pos{i, j}))
		}

		tiles = append(tiles, currentTiles)
	}

	return &Maze{tiles}
}

func (m *Maze) StartingPos() (Pos, error) {
	for i := range m.tiles {
		for j := range m.tiles[i] {
			if m.tiles[i][j].Char() == 'S' {
				return Pos{j, i}, nil
			}
		}
	}

	return Pos{0, 0}, errors.New("No starting position found")
}
