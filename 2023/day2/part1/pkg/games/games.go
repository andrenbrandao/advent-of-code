package games

import (
	"strconv"
	"strings"
)

type Play struct {
	blue  int
	red   int
	green int
}

func NewPlay(input string) *Play {
	cubes := strings.Split(input, ",")
	colorMap := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	for _, cube := range cubes {
		trimmedCube := strings.Trim(cube, " ")
		values := strings.Split(trimmedCube, " ")

		count, color := values[0], values[1]
		colorMap[color], _ = strconv.Atoi(count)
	}

	return &Play{
		blue:  colorMap["blue"],
		red:   colorMap["red"],
		green: colorMap["green"],
	}
}

func (p *Play) Blue() int {
	return p.blue
}

func (p *Play) Red() int {
	return p.red
}

func (p *Play) Green() int {
	return p.red
}

func (p *Play) Valid(validPlay *Play) bool {
	return p.red <= validPlay.red &&
		p.blue <= validPlay.blue &&
		p.green <= validPlay.green
}

type Game struct {
	plays []*Play
}

func NewGame(input string) *Game {
	playsInputString := strings.Split(input, ":")[1]
	playsInput := strings.Split(playsInputString, ";")

	var plays []*Play

	for _, playInput := range playsInput {
		trimmedInput := strings.Trim(playInput, " ")
		plays = append(plays, NewPlay(trimmedInput))
	}

	return &Game{plays}
}

func (g *Game) Plays() []*Play {
	return g.plays
}
