package bot

import (
	"connect4-bot/internal/model"
)

// MyBot implements the Bot interface for an template bot.
type MyBot struct {
}

// Run processes the current game status and determines in which column the coin
// should be inserted.
// In this implementation, the bot simply does nothing and return always a zero.
func (b *MyBot) Run(state *model.StateData) int {

	//
	// Implement your logic here so that the bot can play
	//

	// Currently, the first column is always selected as the next move
	return 0
}

// GetName returns the name of the bot. This name is used for creating the
// WebSocket URL and displaying the bot's name within the game.
func (b *MyBot) GetName() string {
	return "MyBot"
}
