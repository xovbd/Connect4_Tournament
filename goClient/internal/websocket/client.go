package websocket

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"connect4-bot/internal/bot"
	"connect4-bot/internal/model"

	"github.com/gorilla/websocket"
)

// Client represents a client that communicates with the game server via WebSocket.
type Client struct {
	Conn     *websocket.Conn // The WebSocket connection
	Bot      bot.Bot         // The bot that interacts with the game
	Port     int             // The port number to connect to
	ClientId int             // client ID which is given by the gameserver
}

// NewClient initializes a new Client with the provided bot and port
func NewClient(bot bot.Bot, port int) *Client {
	return &Client{
		Bot:  bot,
		Port: port,
	}
}

// Connect establishes the WebSocket connection to the server using the bot's name and the specified port.
func (c *Client) Connect() error {
	// Construct the URL for the WebSocket connection
	url := "ws://localhost:" + strconv.Itoa(c.Port)

	// Establish the WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		// Return error if connection fails
		return err
	}

	// Create the ConnectionRequest
	req := model.ConnectionRequest{
		Type: "connect",
		Name: c.Bot.GetName(),
	}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		log.Fatal("Error while marshalling the ConnectionRequest:", err)
	}

	// Send the ConnectionRequest
	err = conn.WriteMessage(websocket.TextMessage, jsonReq)
	if err != nil {
		log.Fatal("Error while sending the ConnectionRequest:", err)
	}

	// Read the raw message from the WebSocket connection
	_, rawMessage, err := conn.ReadMessage()
	if err != nil {
		// Log and exit if an error occurs while reading the message
		log.Fatal("Error while reading the ConnectionData response:", err)
	}

	// Use a JSON decoder to unmarshal the raw message into the PlayState struct
	decoder := json.NewDecoder(bytes.NewReader(rawMessage))
	var connData model.ConnectionData
	err = decoder.Decode(&connData)
	if err != nil {
		log.Fatal("Error while decoding the ConnectionData response:", string(rawMessage))
	}
	c.ClientId = connData.ID

	// Store the WebSocket connection
	c.Conn = conn
	return nil
}

// Listen requests the current game state and send also the next bot move
func (c *Client) Listen() {
	for {
		req := model.StateRequest{
			Type: "getState",
			ID:   c.ClientId,
		}

		jsonReq, err := json.Marshal(req)
		if err != nil {
			log.Fatal("Error while marshalling the StateRequest:", err)
		}
		log.Println("---> SEND GameStateRequest")
		err = c.Conn.WriteMessage(websocket.TextMessage, jsonReq)
		if err != nil {
			log.Fatal("Error while sending the StateRequest:", err)
		}

		// Read the raw message from the WebSocket connection
		_, rawMessage, err := c.Conn.ReadMessage()
		log.Println("<--- RECEIVE GameStateRequest")
		if err != nil {
			log.Fatal("Error while reading the StateData response:", err)
		}

		// unmarshal the raw message into the StateData struct
		var state model.StateData
		if err := json.Unmarshal(rawMessage, &state); err != nil {
			log.Fatal("Error while unmarshalling the StateData response:", err)
		}

		switch state.GameState {
		case "pending":
			log.Println("     GameState: PENDING")

		case "finished":
			log.Println("     GameState: FINISHED")
			log.Println("     Game finished. Close connection.")
			c.Close()
			return

		case "playing":
			log.Println("     GameState: PLAYING")

			log.Println("---> SEND next move")
			nextMove := model.PlayRequest{
				Column: c.Bot.Run(&state),
				ID:     state.ID,
				Type:   "play",
			}
			jsonReq, err := json.Marshal(nextMove)
			if err != nil {
				log.Fatal("Error while marshalling the PlayRequest:", err)
			}

			err = c.Conn.WriteMessage(websocket.TextMessage, jsonReq)
			if err != nil {
				log.Fatal("Error while sending the PlayRequest:", err)
			}

			_, rawMessage, err := c.Conn.ReadMessage()
			if err != nil {
				log.Fatal("Error while reading the StateData response:", err)
			}
			if err := json.Unmarshal(rawMessage, &state); err != nil {
				log.Fatal("Error while unmarshalling the StateRequest:", err)
			}

			if state.GameState == "finished" {
				log.Println("     GameState: FINISHED")
				log.Println("     Game finished. Close connection.")
				c.Close()
				return
			}

		default:
			log.Println("     GameState: not my turn")
		}

		// Relieve server and client load
		time.Sleep(250 * time.Millisecond)
	}
}

// Close closes the WebSocket connection when done.
func (c *Client) Close() {
	// Close the WebSocket connection
	c.Conn.Close()
}
