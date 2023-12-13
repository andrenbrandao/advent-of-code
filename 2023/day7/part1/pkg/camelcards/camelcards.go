package camelcards

import (
	"errors"
	"sort"
)

type CardType rune

var cardTypes = []CardType{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

type Card struct {
	cardType CardType
}

func NewCard(ct CardType) *Card {
	return &Card{cardType: ct}
}

// StrongerThan returns if one card is stronger than the other
func (c *Card) StrongerThan(other *Card) bool {
	return indexOf(cardTypes, c.cardType) < indexOf(cardTypes, other.cardType)
}

// Equals compares two cards
func (c *Card) Equals(other *Card) bool {
	return c.cardType == other.cardType
}

// indexOf finds the index of the cardType in the cardTypes array
func indexOf(cardTypes []CardType, ct CardType) int {
	for i, lstType := range cardTypes {
		if lstType == ct {
			return i
		}
	}

	return -1
}

// Ordered from most valuable to least valuable
const (
	FIVE_OF_KIND = iota
	FOUR_OF_KIND
	FULL_HOUSE
	THREE_OF_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

type Hand struct {
	cards []*Card
}

func NewHand(hand string) *Hand {
	cards := []*Card{}
	for _, c := range hand {
		cards = append(cards, NewCard(CardType(c)))
	}

	return &Hand{cards}
}

func (h *Hand) Sort() {
	sort.Slice(h.cards, func(i, j int) bool {
		return h.cards[i].StrongerThan(h.cards[j])
	})
}

func (h *Hand) Type() int {
	h.Sort()

	for _, typeChecker := range typeCheckers {
		handType, err := typeChecker(h)
		if err == nil {
			return handType
		}
	}

	return -1
}

var typeCheckers = []func(h *Hand) (int, error){
	fiveOfKindChecker,
	fourOfKindChecker,
	fullHouseChecker,
	threeOfKindChecker,
	twoPairChecker,
	highCardChecker,
}

func fiveOfKindChecker(h *Hand) (int, error) {
	_, maxCount := calculateBuckets(h)

	if maxCount == 5 {
		return FIVE_OF_KIND, nil
	}

	return -1, errors.New("Not of this type")
}

func fourOfKindChecker(h *Hand) (int, error) {
	_, maxCount := calculateBuckets(h)

	if maxCount == 4 {
		return FOUR_OF_KIND, nil
	}

	return -1, errors.New("Not of this type")
}

func fullHouseChecker(h *Hand) (int, error) {
	buckets, _ := calculateBuckets(h)

	if len(buckets) == 2 && buckets[0] == 2 && buckets[1] == 3 {
		return FULL_HOUSE, nil
	}

	return -1, errors.New("Not of this type")
}

func threeOfKindChecker(h *Hand) (int, error) {
	buckets, maxCount := calculateBuckets(h)

	if len(buckets) == 3 && maxCount == 3 {
		return THREE_OF_KIND, nil
	}

	return -1, errors.New("Not of this type")
}

func twoPairChecker(h *Hand) (int, error) {
	buckets, _ := calculateBuckets(h)

	if len(buckets) == 3 && buckets[1] == 2 && buckets[2] == 2 {
		return TWO_PAIR, nil
	}

	return -1, errors.New("Not of this type")
}

func highCardChecker(h *Hand) (int, error) {
	buckets, _ := calculateBuckets(h)

	if len(buckets) == 5 {
		return HIGH_CARD, nil
	}

	return -1, errors.New("Not of this type")
}

func calculateBuckets(h *Hand) ([]int, int) {
	buckets := []int{}

	maxCount := 1
	currentCount := 1

	for i := 1; i < len(h.cards); i++ {
		if h.cards[i].Equals(h.cards[i-1]) {
			currentCount++
		} else {
			buckets = append(buckets, currentCount)
			currentCount = 1
		}

		maxCount = max(maxCount, currentCount)
	}

	buckets = append(buckets, currentCount)

	sort.Slice(buckets, func(i, j int) bool {
		return buckets[i] < buckets[j]
	})

	return buckets, maxCount
}
