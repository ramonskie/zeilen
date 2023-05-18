package game

import (
	"github.com/ramonskie/zeilen/deck"
)

// CalculatePoints calculates the points for a hand of cards according to the rules of the game
func CalculatePoints(hand []deck.Card) int {
	score := 0
	pairs := make(map[string]int)
	for _, card := range hand {
		switch card.Rank {
		case "7":
			score += 7
			pairs[card.Rank]++
			if pairs[card.Rank] == 2 {
				score += 10
			}
		case "8":
			score += 8
			pairs[card.Rank]++
			if pairs[card.Rank] == 2 {
				score += 10
			}
		case "9":
			score += 9
			pairs[card.Rank]++
			if pairs[card.Rank] == 2 {
				score += 10
			}
		case "10":
			score += 10
			pairs[card.Rank]++
			if pairs[card.Rank] == 2 {
				score += 10
			}
		case "jack":
			score += 1
			pairs[card.Rank]++
			if pairs[card.Rank] == 2 {
				score += 30
			}
		case "queen":
			score += 2
			pairs[card.Rank]++
			if pairs[card.Rank] == 2 {
				score += 30
			}
		case "king":
			score += 3
			pairs[card.Rank]++
			if pairs[card.Rank] == 2 {
				score += 30
			}
		case "ace":
			score += 11
			pairs[card.Rank]++
			if pairs[card.Rank] == 2 {
				score += 30
			}
		}

		// Check if there are any special combinations
		if hasKingQueenSameSuit(hand) {
			score += 30
		}
	}
	return score
}

func hasKingQueenSameSuit(hand []deck.Card) bool {
	kingSuit := ""
	queenSuit := ""

	for _, card := range hand {
		if card.Rank == "king" {
			kingSuit = card.Suit
		} else if card.Rank == "queen" {
			queenSuit = card.Suit
		}
	}

	return kingSuit != "" && kingSuit == queenSuit
}
