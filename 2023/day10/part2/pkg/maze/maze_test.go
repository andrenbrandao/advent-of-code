package maze

import (
	"reflect"
	"testing"
)

func TestPipe(t *testing.T) {

	t.Run("returns the next positions", func(t *testing.T) {
		tests := []struct {
			pipeType rune
			want     []Pos
		}{
			{'|', []Pos{{0, -1}, {0, 1}}},
			{'-', []Pos{{-1, 0}, {1, 0}}},
			{'L', []Pos{{0, -1}, {1, 0}}},
			{'J', []Pos{{-1, 0}, {0, -1}}},
			{'7', []Pos{{-1, 0}, {0, 1}}},
			{'F', []Pos{{1, 0}, {0, 1}}},
			{'S', []Pos{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}},
		}

		for _, tt := range tests {
			pipe := NewPipe(tt.pipeType, Pos{0, 0})

			got := pipe.Neighbors()
			want := tt.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("type %c, got %v, want %v", tt.pipeType, got, want)
			}
		}
	})

	t.Run("can only enter through the entrance", func(t *testing.T) {
		left := Pos{-1, 0}
		top := Pos{0, -1}
		right := Pos{1, 0}
		bottom := Pos{0, 1}

		tests := []struct {
			pipeType rune
			lastPos  Pos
			want     bool
		}{
			{'|', left, false},
			{'|', top, true},
			{'|', right, false},
			{'|', bottom, true},
			{'-', left, true},
			{'-', top, false},
			{'-', right, true},
			{'-', bottom, false},
			{'L', left, false},
			{'L', top, true},
			{'L', right, true},
			{'L', bottom, false},
			{'J', left, true},
			{'J', top, true},
			{'J', right, false},
			{'J', bottom, false},
			{'7', left, true},
			{'7', top, false},
			{'7', right, false},
			{'7', bottom, true},
			{'F', left, false},
			{'F', top, false},
			{'F', right, true},
			{'F', bottom, true},
			{'S', left, true},
			{'S', top, true},
			{'S', right, true},
			{'S', bottom, true},
		}

		for _, tt := range tests {
			pipe := NewPipe(tt.pipeType, Pos{0, 0})

			got := pipe.CanEnterFrom(tt.lastPos)
			want := tt.want

			if got != want {
				t.Errorf("type %c, got %v, want %v", tt.pipeType, got, want)
			}
		}
	})
}

func TestGround(t *testing.T) {

	t.Run("returns the next positions as empty array", func(t *testing.T) {
		tests := []struct {
			typ  rune
			want []Pos
		}{
			{'.', []Pos{}},
		}

		for _, tt := range tests {
			ground := NewGround(Pos{0, 0})

			got := ground.Neighbors()
			want := tt.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("type %c, got %v, want %v", tt.typ, got, want)
			}
		}
	})

}

func TestMaze(t *testing.T) {

	t.Run("returns the starting position", func(t *testing.T) {
		tests := []struct {
			maze string
			want Pos
		}{
			{`..S`, Pos{2, 0}},
			{`...
.S.
...`, Pos{1, 1}},
		}

		for _, tt := range tests {
			maze := NewMazeFromString(tt.maze)

			got, _ := maze.StartingPos()
			want := tt.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("maze %s, got %v, want %v", tt.maze, got, want)
			}
		}
	})

	t.Run("counts steps to walk the cycle from the starting position", func(t *testing.T) {
		tests := []struct {
			maze string
			want int
		}{
			{`
.....
.S-7.
.|.|.
.L-J.
.....`, 8},
			{`
.....
.S.7.
.|.|.
.L-J.
.....`, 0},
			{`
.....
.SF7.
.|.|.
.L-J.
.....`, 0}, // should only walk through entrance and exit of pipes
			{`
-L|F7
7S-7|
L|7||
-L-J|
L|-JF`, 8},
		}

		for _, tt := range tests {
			maze := NewMazeFromString(tt.maze)

			got := maze.CycleSteps()
			want := tt.want

			if got != want {
				t.Errorf("maze %s, got %v, want %v", tt.maze, got, want)
			}
		}
	})

	t.Run("farthest steps from the starting position", func(t *testing.T) {
		tests := []struct {
			maze string
			want int
		}{
			{`
.....
.S-7.
.|.|.
.L-J.
.....`, 4},
			{`
..F7.
.FJ|.
SJ.L7
|F--J
LJ...`, 8},
		}

		for _, tt := range tests {
			maze := NewMazeFromString(tt.maze)

			got := maze.FarthestSteps()
			want := tt.want

			if got != want {
				t.Errorf("maze %s, got %v, want %v", tt.maze, got, want)
			}
		}
	})

	t.Run("count the number of enclosed tiles by the loop", func(t *testing.T) {
		tests := []struct {
			maze string
			want int
		}{
			{`
.....
.S-7.
.|.|.
.L-J.
.....`, 1},
			{`
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`, 4},
			{`
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`, 8},
			{`
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`, 10},
		}

		for _, tt := range tests {
			maze := NewMazeFromString(tt.maze)

			got := maze.NumberEnclosedTiles()
			want := tt.want

			if got != want {
				t.Errorf("maze %s, got %v, want %v", tt.maze, got, want)
			}
		}
	})
}
