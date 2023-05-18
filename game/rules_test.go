package game_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ramonskie/zeilen/deck"
	"github.com/ramonskie/zeilen/game"
)

var _ = Describe("CalculatePoints", func() {
	Context("when calculating points for a hand of cards", func() {
		It("should correctly calculate the points for a single card", func() {
			hand := []deck.Card{{Rank: "7", Suit: "hearts"}} // a player should never have a single card
			expectedPoints := 7
			Expect(game.CalculatePoints(hand)).To(Equal(expectedPoints))
		})

		It("should correctly calculate the points for two different cards", func() {
			hand := []deck.Card{{Rank: "7", Suit: "hearts"}, {Rank: "king", Suit: "diamonds"}}
			expectedPoints := 10
			Expect(game.CalculatePoints(hand)).To(Equal(expectedPoints))
		})

		It("should correctly calculate the points for two of the same number cards", func() {
			hand := []deck.Card{{Rank: "10", Suit: "clubs"}, {Rank: "10", Suit: "hearts"}}
			expectedPoints := 30 // jack=10 + jack=10 + 10
			Expect(game.CalculatePoints(hand)).To(Equal(expectedPoints))
		})

		It("should correctly calculate the points for two of the same face cards", func() {
			hand := []deck.Card{{Rank: "jack", Suit: "clubs"}, {Rank: "jack", Suit: "hearts"}}
			expectedPoints := 32 // jack=1 + jack=1 + 30
			Expect(game.CalculatePoints(hand)).To(Equal(expectedPoints))
		})

		It("should correctly calculate the points for two aces", func() {
			hand := []deck.Card{{Rank: "ace", Suit: "clubs"}, {Rank: "ace", Suit: "hearts"}}
			expectedPoints := 52 //ace=11 + ace=11 + 30
			Expect(game.CalculatePoints(hand)).To(Equal(expectedPoints))
		})

		It("should correctly calculate the points for king and queen of the different suit", func() {
			hand := []deck.Card{{Rank: "queen", Suit: "spades"}, {Rank: "king", Suit: "hearts"}}
			expectedPoints := 5
			Expect(game.CalculatePoints(hand)).To(Equal(expectedPoints))
		})

		// highest possible set of cards
		It("should correctly calculate the points for king and queen of the same suit", func() {
			hand := []deck.Card{{Rank: "queen", Suit: "spades"}, {Rank: "king", Suit: "spades"}}
			expectedPoints := 65
			Expect(game.CalculatePoints(hand)).To(Equal(expectedPoints))
		})

		// lowest possible set of cards
		It("should correctly calculate the points for two different cards", func() {
			hand := []deck.Card{{Rank: "jack", Suit: "hearts"}, {Rank: "queen", Suit: "diamonds"}}
			expectedPoints := 3
			Expect(game.CalculatePoints(hand)).To(Equal(expectedPoints))
		})
	})
})
