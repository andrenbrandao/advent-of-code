package camelcards

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type HandBid struct {
	hand *Hand
	bid  int
}
type Game struct {
	hands    []*Hand
	handBids []*HandBid
}

func NewGame(handStrings []string, bids []string) *Game {
	hands := []*Hand{}
	handBids := []*HandBid{}

	for i, hs := range handStrings {
		hand := NewHand(hs)
		hands = append(hands, hand)

		bid, _ := strconv.Atoi(bids[i])
		handBids = append(handBids, &HandBid{hand, bid})
	}

	return &Game{hands, handBids}
}

func NewGameFromInput(input string) *Game {
	lines := strings.Split(input, "\n")

	hands := []string{}
	bids := []string{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		hands = append(hands, fields[0])
		bids = append(bids, fields[1])
	}

	game := NewGame(hands, bids)
	return game
}

// Sort hands from weakest to strongest
func (g *Game) Sort() {
	sort.Slice(g.hands, func(i, j int) bool {
		return !g.hands[i].StrongerThan(g.hands[j])
	})

	sort.Slice(g.handBids, func(i, j int) bool {
		return !g.handBids[i].hand.StrongerThan(g.handBids[j].hand)
	})
}

func (g *Game) String() string {
	return fmt.Sprintf("%s", g.hands)
}

func (g *Game) Winnings() int {
	g.Sort()

	total := 0
	for i, handbid := range g.handBids {
		bid := handbid.bid

		total += bid * (i + 1)
	}

	return total
}
