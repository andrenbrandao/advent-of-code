package camelcards

import (
	"errors"
	"fmt"
	"sort"
)

// Ordered from least valuable to most valuable
const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_KIND
	FULL_HOUSE
	FOUR_OF_KIND
	FIVE_OF_KIND
)

type Hand struct {
	cards []*Card
}

func NewHand(hand string) *Hand {
	cards := []*Card{}
	for _, c := range hand {
		cards = append(cards, NewCard(CardType(c)))
	}

	return NewHandWithCards(cards)
}

func NewHandWithCards(cards []*Card) *Hand {
	return &Hand{cards}
}

func (h *Hand) StrongerThan(other *Hand) bool {
	type1 := h.Type()
	type2 := other.Type()

	if type1 != type2 {
		return type1 > type2
	}

	for i := 0; i < len(h.cards); i++ {
		if h.cards[i].Equals(other.cards[i]) {
			continue
		}

		return h.cards[i].StrongerThan(other.cards[i])
	}

	panic(errors.New("Internal error"))
}

func (h *Hand) String() string {
	return fmt.Sprintf("%s", h.cards)
}

// Sort cards from strongest to weakest
func (h *Hand) Sort() {
	sort.Slice(h.cards, func(i, j int) bool {
		return h.cards[i].StrongerThan(h.cards[j])
	})
}

func (h *Hand) Type() int {
	hands := h.generateHandsOptimized()

	types := []int{}
	for _, hand := range hands {
		typ := hand.calculateType()
		types = append(types, typ)
	}

	sort.Slice(types, func(i, j int) bool {
		return types[i] > types[j]
	})

	return types[0]
}

func (h *Hand) calculateType() int {
	for _, typeChecker := range typeCheckers {
		handType, err := typeChecker(h)
		if err == nil {
			return handType
		}
	}

	return -1
}

// generateHands creates all possible hands when we have a joker
func (h *Hand) generateHands() []*Hand {
	hands := []*Hand{h}

	var recHelper func(pos int, currentCards []*Card)
	recHelper = func(pos int, currentCards []*Card) {
		if pos == len(h.cards) {
			copiedCards := make([]*Card, len(h.cards))
			copy(copiedCards, currentCards)
			hands = append(hands, NewHandWithCards(copiedCards))
			return
		}

		if h.cards[pos].IsJoker() {
			for _, typ := range cardTypes {
				c := NewCard(typ)
				currentCards = append(currentCards, c)
				recHelper(pos+1, currentCards)
				currentCards = currentCards[:len(currentCards)-1]
			}
		} else {
			currentCards = append(currentCards, h.cards[pos])
			recHelper(pos+1, currentCards)
			currentCards = currentCards[:len(currentCards)-1]
		}
	}

	recHelper(0, []*Card{})
	return hands
}

func (h *Hand) generateHandsOptimized() []*Hand {
	hands := []*Hand{}
	possibleTypesSet := map[CardType]bool{}

	var recHelper func(pos int, currentCards []*Card, possibleTypes []CardType)
	recHelper = func(pos int, currentCards []*Card, possibleTypes []CardType) {
		if pos == len(h.cards) {
			copiedCards := make([]*Card, len(h.cards))
			copy(copiedCards, currentCards)
			hands = append(hands, NewHandWithCards(copiedCards))
			return
		}

		if h.cards[pos].IsJoker() {
			for _, typ := range possibleTypes {
				c := NewCard(typ)
				currentCards = append(currentCards, c)
				recHelper(pos+1, currentCards, possibleTypes)
				currentCards = currentCards[:len(currentCards)-1]
			}
		} else {
			currentCards = append(currentCards, h.cards[pos])
			recHelper(pos+1, currentCards, possibleTypes)
			currentCards = currentCards[:len(currentCards)-1]
		}
	}

	for _, c := range h.cards {
		if !c.IsJoker() {
			possibleTypesSet[c.cardType] = true
		}
	}

	keys := make([]CardType, 0, len(possibleTypesSet))
	for ct := range possibleTypesSet {
		keys = append(keys, ct)
	}

	if len(keys) == 0 {
		keys = []CardType{'A'}
	}

	recHelper(0, []*Card{}, keys)
	return hands
}

var typeCheckers = []func(h *Hand) (int, error){
	fiveOfKindChecker,
	fourOfKindChecker,
	fullHouseChecker,
	threeOfKindChecker,
	twoPairChecker,
	onePairChecker,
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

func onePairChecker(h *Hand) (int, error) {
	buckets, maxCount := calculateBuckets(h)

	if len(buckets) == 4 && maxCount == 2 {
		return ONE_PAIR, nil
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
	cards := make([]*Card, len(h.cards))
	copy(cards, h.cards)
	hand := NewHandWithCards(cards)
	hand.Sort()

	buckets := []int{}

	maxCount := 1
	currentCount := 1

	for i := 1; i < len(hand.cards); i++ {
		if hand.cards[i].Equals(hand.cards[i-1]) {
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
