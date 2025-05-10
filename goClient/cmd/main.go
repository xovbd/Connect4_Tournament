package main

import (
	"flag"
	"fmt"
	"log"

	"connect4-bot/internal/bot"
	"connect4-bot/internal/websocket"
)

func main() {
	// Define command-line arguments
	port := flag.Int("port", 8765, "Port on which the game server listens")
	botType := flag.String("bot", "MyBot", "Name of the bot to be used")

	// Parse the command-line arguments
	flag.Parse()

	// Print the port and selected bot information
	fmt.Printf("Server listens on port: %d\n", *port)
	fmt.Printf("Selected bot          : %s\n", *botType)

	// Use the BotFactory to create the desired bot
	factory := &bot.BotFactory{}
	myBot, err := factory.NewBot(*botType)
	if err != nil {
		log.Fatal("Error creating the bot:", err)
	}

	// Create and connect the WebSocket client
	client := websocket.NewClient(myBot, *port)

	err = client.Connect()
	if err != nil {
		log.Fatal("Connection failed:", err)
	} else {
		log.Println("Client ID             :", client.ClientId)
	}

	// check if we have a valid connection
	if client.ClientId == 0 {
		log.Fatal("Connection failed.")
	}

	// Start goroutines
	go client.Listen() // Listen for WebSocket messages and call the bot

	// Block the main thread to keep the program running
	select {}
}
