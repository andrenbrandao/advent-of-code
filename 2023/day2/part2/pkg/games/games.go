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

func (p *Play) MinPlay(other *Play) *Play {
	blue := max(p.blue, other.blue)
	red := max(p.red, other.red)
	green := max(p.green, other.green)
	return &Play{blue, red, green}
}

func (p *Play) PowerSet() int {
	return p.blue * p.red * p.green
}

type Game struct {
	id    int
	plays []*Play
}

func NewGame(input string) *Game {
	inputSplitted := strings.Split(input, ":")
	gameString := inputSplitted[0]
	id := strings.Split(gameString, " ")[1]
	playsInputString := inputSplitted[1]
	playsInput := strings.Split(playsInputString, ";")

	var plays []*Play

	for _, playInput := range playsInput {
		trimmedInput := strings.Trim(playInput, " ")
		plays = append(plays, NewPlay(trimmedInput))
	}

	idInt, _ := strconv.Atoi(string(id))
	return &Game{idInt, plays}
}

func (g *Game) Plays() []*Play {
	return g.plays
}

func (g *Game) Id() int {
	return g.id
}

func (g *Game) Valid(validPlay *Play) bool {
	for _, play := range g.plays {
		if !play.Valid(validPlay) {
			return false
		}
	}

	return true
}

func (g *Game) MinPlay() *Play {
	result := g.plays[0]

	for _, play := range g.plays[1:] {
		result = result.MinPlay(play)
	}

	return result
}

type GameRecords struct {
	games []*Game
}

func NewGameRecords(gamesInput []string) *GameRecords {
	games := make([]*Game, len(gamesInput))
	for i, gameInput := range gamesInput {
		games[i] = NewGame(gameInput)
	}

	return &GameRecords{games}
}

func (gr *GameRecords) ValidGames(validPlay *Play) []*Game {
	var validGames []*Game

	for _, game := range gr.games {
		if game.Valid(validPlay) {
			validGames = append(validGames, game)
		}
	}

	return validGames
}

func (gr *GameRecords) SumValidGameIds(validPlay *Play) int {
	sum := 0

	for _, game := range gr.games {
		if game.Valid(validPlay) {
			sum += game.Id()
		}
	}

	return sum
}

func (gr *GameRecords) SumPowerSetCubes() int {
	sum := 0

	for _, game := range gr.games {
		minPlay := game.MinPlay()
		sum += minPlay.PowerSet()
	}

	return sum
}
