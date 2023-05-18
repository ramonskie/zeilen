package deck

import (
	"math/rand"
	"time"
)

// Card represents a playing card with a rank and suit
type Card struct {
	Rank string
	Suit string
}

// Deck represents a collection of cards
type Deck struct {
	Cards []Card
}

// NewDeck creates a new deck of cards in sorted order
func NewDeck() *Deck {
	var ranks = []string{"7", "8", "9", "10", "jack", "queen", "king", "ace"}
	var suits = []string{"hearts", "diamonds", "clubs", "spades"}

	deck := &Deck{Cards: []Card{}}
	for _, rank := range ranks {
		for _, suit := range suits {
			card := Card{Rank: rank, Suit: suit}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

// Shuffle shuffles the deck of cards randomly
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// DrawCards removes n cards from the deck and returns them
func (d *Deck) DrawCards(n int) []Card {
	cards := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return cards
}
