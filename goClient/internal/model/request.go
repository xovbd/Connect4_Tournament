package model

// ConnectionRequest represents a request to connect with the gameserver
type ConnectionRequest struct {
	Type string `json:"type"`
	Name string `json:"name"` // Request type: always "connect"
}

// PlayRequest represents a request to make a move in a specific column.
type PlayRequest struct {
	ID     int    `json:"id"`     // Unique identifier of the bot or client
	Type   string `json:"type"`   // Request type: always "play"
	Column int    `json:"column"` // Column index where the move should be played
}

// StateRequest represents a request to retrieve the current game state.
type StateRequest struct {
	ID   int    `json:"id"`   // Unique identifier of the bot or client
	Type string `json:"type"` // Request type: always "getState"
}
