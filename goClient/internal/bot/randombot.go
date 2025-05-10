package bot

import (
	"connect4-bot/internal/model"
	"math/rand"
	"time"
)

// RandomBot implements the Bot interface for an random playing bot.
type RandomBot struct {
}

// Run processes the current game status and determines in which column the coin
// should be inserted.
// In this implementation, the bot simply insert the coin in random columns.
func (b *RandomBot) Run(state *model.StateData) int {
	// generate a seed for the random generator
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// return a random column number
	return rand.Intn(len(state.Field[0]))
}

// GetName returns the name of the bot. This name is used for creating the
// WebSocket URL and displaying the bot's name within the game.
func (b *RandomBot) GetName() string {
	return "RandomBot"
}
