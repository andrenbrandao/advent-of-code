package maze

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

type Ground struct {
	pos Pos
}

func NewGround(pos Pos) *Ground {
	return &Ground{pos}
}

func (g *Ground) Neighbors() []Pos {
	return []Pos{}
}
