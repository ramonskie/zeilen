package game

import (
	"fmt"

	"github.com/ramonskie/zeilen/deck"

	"github.com/gorilla/websocket"
)

// Player represents a player in the game with a hand of cards and a balance of credits
type Player struct {
	ID       string
	Hand     []deck.Card
	Credit   int
	Balance  int
	Folded   bool
	Conn     *websocket.Conn
	GameRoom *GameRoom
}

// GameRoom represents a game room where players can join and play together
type GameRoom struct {
	ID      string
	Players []*Player
}

// Game represents a game of cards with a deck, a pot of credits, and a list of players
type Game struct {
	Deck    *deck.Deck
	Pot     int
	Players []*Player
}

// NewGame creates a new game with the specified number of players
func NewGame(numPlayers int) *Game {
	deck := deck.NewDeck()
	deck.Shuffle()

	players := make([]*Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = &Player{
			ID:      "",
			Credit:  1,
			Balance: 1,
		}
	}

	return &Game{
		Deck:    deck,
		Pot:     0,
		Players: players,
	}
}

// initializeGame initializes the game with a new deck and all players adding their credits to the pot
func InitializeGame(numPlayers int) *Game {
	game := NewGame(numPlayers)
	for _, player := range game.Players {
		game.Pot += player.Credit
		player.Balance -= player.Credit
	}
	return game
}

// isGameOver checks if the game is over by checking if any player has a balance greater than 0
func IsGameOver(game *Game) bool {
	for _, player := range game.Players {
		if player.Balance == 0 {
			return false
		}
	}
	return true
}

// DetermineWinner determines the winner of the game and returns the winning player and their point total
func DetermineWinner(players []*Player) (*Player, int) {
	var winner *Player
	highestPoints := 0

	// Calculate each player's points
	for _, player := range players {
		if player.Folded {
			continue
		}

		points := CalculatePoints(player.Hand)
		fmt.Printf("Player %s has %d points\n", player.ID, points)

		if points > highestPoints {
			// New highest score, reset winners
			winner = player
			highestPoints = points
		} else if points == highestPoints {
			// Same score as current highest both player loose
			winner = nil
		}

		// if someone wins with jack and queen when everyone folded they win and get from every player a the amount in the pot
	}

	return winner, highestPoints
}

var unregister = make(chan *Player)

func (p *Player) Listen() {
	defer func() {
		unregister <- p
		p.GameRoom.Broadcast([]byte(p.ID + " left the game room"))
		p.Conn.Close()
	}()

	for {
		_, message, err := p.Conn.ReadMessage()
		if err != nil {
			break
		}

		p.GameRoom.Broadcast([]byte(p.ID + ": " + string(message)))
	}
}

func (p *Player) Send(message []byte) {
	err := p.Conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return
	}
}

func (gr *GameRoom) Broadcast(message []byte) {
	for _, player := range gr.Players {
		player.Send(message)
	}
}