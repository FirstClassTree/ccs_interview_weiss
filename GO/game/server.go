package game

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Game struct {
	mu            sync.Mutex
	players       [2]net.Conn // Two players
	turnIndex     int         // 0 or 1 (Whose turn it is)
	correctNumber int         // Number to be guessed
	gameOver      bool        // Flag to track if game is over
}

var game = Game{
	turnIndex: 0,
	gameOver:  false,
}

var workQueue = make(chan int, 2)

func StartServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server started, waiting for players...")
	// Accept two players
	for i := 0; i < 2; i++ {
		conn, err := listener.Accept() // This line blocks until a player connects
		if err != nil {
			log.Fatalf("Error accepting connection: %v", err)
		}
		game.players[i] = conn
		fmt.Printf("Player %d connected.\n", i+1)
		writeToClient(conn, fmt.Sprintf("You are Player %d. Waiting for the game to start...\n", i+1))
	}

	// Generate a random correct number
	var one int = 1
	var two int = 2
	game.correctNumber = generateCorrectNumber(&one, &two)
	fmt.Println("Secret number generated.")

	// Notify players that the game has started
	broadcastMessage("🎮 Game started! Player 1 goes first.")

	// **Start game loop in a goroutine**
	go handleTurn()

	// Prevent server from exiting immediately
	select {} // Blocks forever to keep the server running
}

// Handles one turn in the game loop
func handleTurn() {
	for !game.gameOver {
		currentPlayer := game.players[game.turnIndex]
		writeToClient(currentPlayer, fmt.Sprintf("It's your turn, Player %d. Enter a number (1-100): ", game.turnIndex+1))
		for {
			// Read input from the current player
			buffer := make([]byte, 1024)
			n, err := currentPlayer.Read(buffer)
			if err != nil {
				log.Printf("Error reading from Player %d: %v", game.turnIndex+1, err)
				handleDisconnection(game.turnIndex)
				return
			}

			// Convert input to string
			guess := string(buffer[:n])
			broadcastMessage(fmt.Sprintf("Player %d guessed: %s\n", game.turnIndex+1, guess))

			// Validate the guess is correct format
			numGuess, err := ValidateGuess(guess)
			if err != nil {
				writeToClient(currentPlayer, "❌ Invalid guess: "+err.Error())
				continue
			}

			// ✅ Check if the guess is correct
			if numGuess == game.correctNumber {
				game.gameOver = true
				winnerMsg := fmt.Sprintf("🎉 Player %d wins! The correct number was %d.", game.turnIndex+1, game.correctNumber)
				fmt.Println(winnerMsg)
				broadcastMessage(winnerMsg)
				return
			}

			// ✅ Switch turn after a valid move
			game.turnIndex = (game.turnIndex + 1) % 2

			// ✅ Notify next player
			broadcastMessage(fmt.Sprintf("🔄 Player %d, it's your turn now!", game.turnIndex+1))
			break // Exit the loop and proceed to the next turn
		}
	}
}

// Handles player disconnection
func handleDisconnection(playerIndex int) {
	if game.players[playerIndex] != nil {
		game.players[playerIndex].Close()
		game.players[playerIndex] = nil
	}
	broadcastMessage(fmt.Sprintf("Player %d has disconnected. Game over.", playerIndex+1))
	game.gameOver = true
}

// Sends a message to both players
func broadcastMessage(message string) {
	for _, conn := range game.players {
		if conn != nil {
			writeToClient(conn, message)
		}
	}
}

func writeToClient(conn net.Conn, s string) {
	_, err := conn.Write([]byte(s))
	if err != nil {
		log.Printf("Error writing to client: %v", err)
	}
}
