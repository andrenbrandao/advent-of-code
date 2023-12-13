package camelcards

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
