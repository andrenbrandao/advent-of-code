
# Algorithm

Camel Cards is a game similar to Poker. In Camel Cards, you get a list of hands, and your goal is to order them based on the strength of each hand.

## Introduction

Hand: consists of 5 cards
Card Types: A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2

Possible Hands:

- Five of a kind, where all five cards have the same label: AAAAA
- Four of a kind, where four cards have the same label and one card has a different label: AA8AA
- Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
- Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
- Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
- One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
- High card, where all cards' labels are distinct: 23456

These hands are ordered from strongest to weakest.

If two hands are of the same type, they are compared by their first cards: 33332 and 2AAAA are of the same type (Three of a kind), but 3 is higher than 2.

## Use Cases

### Card

- Can be compared to another card to find which one is stronger

### Hand

- Can be compare to another hadn to find which one is stronger
- Has a type calculated based on the cards

### Game

- Receives a list of hands with bids
- Returns the winnings by sorting the hands and multiplying the rank of each hand by the bid and summing them

Now, how do we sort the hands? We can use a sorting algorithm like quicksort, but comparing
the types of each hand. Use the method strongerThan(hand) as a comparison algorithm.

## Design

```
CardType:
- A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2

HandType:
- Five of a kind
- Four of a kind
- Full house
- Three of a kind
- Two pair
- One Pair
- High card

Card:
- type
--
+ strongerThan(card Card)

Hand:
- cards []Card
--
+ strongerThan(hand Hand)
+ type() HandType

Bid: int

Game:
- hands []Hand
- bids []Bid
--
- sortHands()
+ Winnings()

```

## Playing the Game

```
Input:

- Hand
- Bid

32T3K 765
T55J5 684
KK677 2
KTJJT 220
QQQJA 483
```

The Game should receive the input and create the Hands and Bids.
We call `Winnings` and should receive the result.
