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

func (p *Pipe) CanEnterFrom(lastPos Pos) bool {
	top := Pos{p.pos.x, p.pos.y - 1}
	bottom := Pos{p.pos.x, p.pos.y + 1}
	left := Pos{p.pos.x - 1, p.pos.y}
	right := Pos{p.pos.x + 1, p.pos.y}

	switch p.pipeType {
	case '|':
		return lastPos == top || lastPos == bottom
	case '-':
		return lastPos == left || lastPos == right
	case 'L':
		return lastPos == top || lastPos == right
	case 'J':
		return lastPos == left || lastPos == top
	case '7':
		return lastPos == left || lastPos == bottom
	case 'F':
		return lastPos == right || lastPos == bottom
	case 'S':
		return lastPos == left || lastPos == top || lastPos == right || lastPos == bottom
	}
	return false
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

func (g *Ground) CanEnterFrom(lastPos Pos) bool {
	return false
}

type Tile interface {
	Neighbors() []Pos
	Char() rune
	CanEnterFrom(p Pos) bool
}
type Maze struct {
	tiles [][]Tile
}

func NewMazeFromString(input string) *Maze {
	tiles := [][]Tile{}
	lines := strings.Split(input, "\n")

	i := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		currentTiles := []Tile{}

		for j, c := range line {
			if c == '.' {
				currentTiles = append(currentTiles, NewGround(Pos{j, i}))
				continue
			}

			currentTiles = append(currentTiles, NewPipe(c, Pos{j, i}))
		}

		tiles = append(tiles, currentTiles)
		i++
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

func (m *Maze) CycleSteps() int {
	visited := map[Pos]bool{}

	startPos, err := m.StartingPos()
	if err != nil {
		panic(err)
	}

	cycleDetected := false

	height := len(m.tiles)
	width := len(m.tiles[0])

	var dfs func(pos Pos, lastPos Pos) (int, error)
	dfs = func(pos Pos, lastPos Pos) (int, error) {

		currentTile := m.tiles[pos.y][pos.x]
		visited[pos] = true

		steps := 0
		for _, neighbor := range currentTile.Neighbors() {
			if neighbor == startPos && startPos != lastPos {
				cycleDetected = true
				return 1, nil
			}

			i := neighbor.y
			j := neighbor.x
			// out of bounds neighbor
			// visited
			// or invalid pipe entrance
			// should be ignored
			if i < 0 || i >= height || j < 0 || j >= width || visited[neighbor] || !m.tiles[neighbor.y][neighbor.x].CanEnterFrom(pos) {
				continue
			}

			currentPathSteps, err := dfs(neighbor, pos)
			if err != nil {
				continue
			}

			steps = max(steps, 1+currentPathSteps)
		}

		return steps, nil
	}

	steps, err := dfs(startPos, startPos)

	if err != nil {
		panic(err)
	}

	if cycleDetected {
		return steps
	}

	return 0
}

func (m *Maze) FarthestSteps() int {
	return m.CycleSteps() / 2
}
