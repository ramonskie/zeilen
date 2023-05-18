package game_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ramonskie/zeilen/deck"
	"github.com/ramonskie/zeilen/game"
)

var _ = Describe("DetermineWinner", func() {
	Context("when determining the winner", func() {
		It("should correctly determine the winner with the highest points", func() {
			player1 := &game.Player{ID: "1", Hand: []deck.Card{{Rank: "7", Suit: "hearts"}, {Rank: "7", Suit: "clubs"}}}
			player2 := &game.Player{ID: "2", Hand: []deck.Card{{Rank: "king", Suit: "diamonds"}, {Rank: "queen", Suit: "spades"}}} //low card should not win
			player3 := &game.Player{ID: "3", Hand: []deck.Card{{Rank: "10", Suit: "hearts"}, {Rank: "ace", Suit: "diamonds"}}}

			players := []*game.Player{player1, player2, player3}

			winner, points := game.DetermineWinner(players)

			Expect(winner).To(Equal(player1))
			Expect(points).To(Equal(24))
		})

		It("should correctly handle multiple players with the same points", func() {
			player1 := &game.Player{ID: "1", Hand: []deck.Card{{Rank: "10", Suit: "hearts"}, {Rank: "10", Suit: "clubs"}}}
			player2 := &game.Player{ID: "2", Hand: []deck.Card{{Rank: "10", Suit: "diamonds"}, {Rank: "10", Suit: "spades"}}}
			player3 := &game.Player{ID: "3", Hand: []deck.Card{{Rank: "8", Suit: "hearts"}, {Rank: "king", Suit: "diamonds"}}}

			players := []*game.Player{player1, player2, player3}

			winner, points := game.DetermineWinner(players)

			Expect(winner).To(BeNil())
			Expect(points).To(Equal(30))
		})

		It("should correctly handle all players folding", func() {
			player1 := &game.Player{ID: "1", Hand: []deck.Card{{Rank: "7", Suit: "hearts"}, {Rank: "7", Suit: "clubs"}}, Folded: true}
			player2 := &game.Player{ID: "2", Hand: []deck.Card{{Rank: "king", Suit: "diamonds"}, {Rank: "queen", Suit: "spades"}}, Folded: true}
			player3 := &game.Player{ID: "3", Hand: []deck.Card{{Rank: "10", Suit: "hearts"}, {Rank: "ace", Suit: "diamonds"}}, Folded: true}

			players := []*game.Player{player1, player2, player3}

			winner, points := game.DetermineWinner(players)

			Expect(winner).To(BeNil())
			Expect(points).To(Equal(0))
		})

		It("should correctly handle player folding with highest score card", func() {
			player1 := &game.Player{ID: "1", Hand: []deck.Card{{Rank: "7", Suit: "hearts"}, {Rank: "7", Suit: "clubs"}}, Folded: true}
			player2 := &game.Player{ID: "2", Hand: []deck.Card{{Rank: "king", Suit: "diamonds"}, {Rank: "queen", Suit: "spades"}}}
			player3 := &game.Player{ID: "3", Hand: []deck.Card{{Rank: "10", Suit: "hearts"}, {Rank: "ace", Suit: "diamonds"}}}

			players := []*game.Player{player1, player2, player3}

			winner, points := game.DetermineWinner(players)

			Expect(winner).To(Equal(player3))
			Expect(points).To(Equal(21))
		})
	})
})
