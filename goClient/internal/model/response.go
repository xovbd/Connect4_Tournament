package model

import "encoding/json"

// ConnectionData represents the response from the server regarding the connection status.
type ConnectionData struct {
	ID        int  `json:"id"`        // Unique identifier for the bot or client.
	Connected bool `json:"connected"` // Indicates whether the connection is established.
}

// StateData represents the game state information returned by the server.
type StateData struct {
	ID        int     `json:"id"`                  // Unique identifier for the bot or client.
	GameState string  `json:"gameState,omitempty"` // The current game state (e.g. "playing", "finished"...)
	Field     [][]int `json:"field,omitempty"`     // The current game field as a 2D array.
}

// own UnmarshalJSON method implementation because the field numbers from the
// gameserver are float, but we want to use integer
func (s *StateData) UnmarshalJSON(data []byte) error {

	var raw struct {
		ID        int             `json:"id"`
		GameState string          `json:"gameState,omitempty"`
		Field     [][]interface{} `json:"field,omitempty"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var field [][]int
	for _, row := range raw.Field {
		intRow := make([]int, len(row))
		for i, val := range row {
			if num, ok := val.(float64); ok {
				intRow[i] = int(num)
			} else {
				intRow[i] = 0
			}
		}
		field = append(field, intRow)
	}

	s.ID = raw.ID
	s.GameState = raw.GameState
	s.Field = field

	return nil
}
