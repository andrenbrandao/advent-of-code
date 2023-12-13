package camelcards

import (
	"fmt"
	"sort"
)

type Game struct {
	hands []*Hand
}

func NewGame(handStrings []string) *Game {
	hands := []*Hand{}
	for _, hs := range handStrings {
		hands = append(hands, NewHand(hs))
	}

	return &Game{hands}
}

// Sort hands from weakest to strongest
func (g *Game) Sort() {
	sort.Slice(g.hands, func(i, j int) bool {
		return g.hands[i].StrongerThan(g.hands[j])
	})
}

func (g *Game) String() string {
	return fmt.Sprintf("%s", g.hands)
}
