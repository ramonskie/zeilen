package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/ramonskie/zeilen/game"
)

var (
	gameRooms map[string]*game.GameRoom
	players   map[string]*game.Player
	upgrader  = websocket.Upgrader{}
	broadcast = make(chan []byte)
	register  = make(chan *game.Player)
	// unregister = make(chan *game.Player) //moved to game/game.go
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")   // Load HTML templates
	router.Static("/static", "./static") // Serve static files

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// 	// Initialize the game
	mygame := game.InitializeGame(4)

	router.POST("/play", func(c *gin.Context) {
		// Deal each player two cards
		for _, player := range mygame.Players {
			player.Hand = mygame.Deck.DrawCards(2)
			fmt.Println("DEBUG: player.Hand:")
		}

		// Render the play.html template and pass the player's ID and hand
		// Prompt each player to play or fold
		for _, player := range mygame.Players {
			// Prompt the player to play or fold
			var action string
			for action != "play" && action != "fold" {
				// Render the play.html template and pass the player's ID and hand
				c.HTML(http.StatusOK, "play.html", gin.H{
					"PlayerID": player.ID,
					"Hand":     player.Hand,
				})

				// Wait for the player's response
				action = getPlayerAction(c)
				fmt.Println("DEBUG: player action:")
			}

			if action == "fold" {
				// If the player folds, reveal their hand and subtract the pot from their balance
				player.Balance -= mygame.Pot
				mygame.Pot = 0
				// player who folds should be removed from winner calculation
				player.Folded = true
				fmt.Println("DEBUG: player folded:")
			} else {
				// If the player plays, keep their hand hidden and add their credit to the pot
				mygame.Pot += player.Credit
				player.Credit = 0
				fmt.Println("DEBUG: player did not fold:")
			}
		}
	})

	router.GET("/ws/:roomID/:playerID", func(c *gin.Context) {
		roomID := c.Param("roomID")
		playerID := c.Param("playerID")

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection to WebSocket"})
			return
		}

		player := &game.Player{
			ID:   playerID,
			Conn: conn,
		}

		register <- player

		// Create or join the game room
		gameRoom, exists := gameRooms[roomID]
		if !exists {
			gameRoom = &game.GameRoom{
				ID:      roomID,
				Players: []*game.Player{player},
			}
			gameRooms[roomID] = gameRoom
		} else {
			gameRoom.Players = append(gameRoom.Players, player)
		}

		player.GameRoom = gameRoom

		// Start listening for messages from the player's WebSocket connection
		go player.Listen()

		// Send a welcome message to the player
		player.Send([]byte("Welcome to the game room!"))

		// Broadcast the player's presence to all other players in the game room
		gameRoom.Broadcast([]byte(playerID + " joined the game room"))
	})

	router.Run(":8080")
}

func getPlayerAction(c *gin.Context) string {
	// Retrieve the player's action from the request body
	action := c.PostForm("action")
	fmt.Println("DEBUG: player action:")
	return action
}
