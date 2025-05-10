# connect4 Tournament - Go Client

## Overview

This is a Go client for the connect4 game.

## Installation

### Prerequisites

- Go 1.23.5 or higher
- github.com/gorilla/websocket v1.5.3 or higher

### Steps to Install

1. **Install Dependencies**

   Make sure to install the necessary Go dependencies. You can do this by
   running:

   ```bash
   go mod tidy
   ```

2. **Build the Bot**

   To compile the bot, simply run:

   ```bash
   (on Linux)
   go build -o connect4-bot cmd/main.go

   (on Windows)
   go build -o connect4-bot.exe cmd/main.go
   ```

3. **Start the Bot**

   You can now start the bot with the following command:

   ```bash
   (on Linux)
   ./connect4-bot

   (on Windows)
   connect4-bot.exe
   ```

   **Configuration**

   You can configure the bot by passing arguments when starting the bot:

   ```bash
   (on Linux)
   ./connect4-bot --port 8765 --bot RandomBot

   (on Windows)
   connect4-bot.exe --port 8765 --bot RandomBot
   ```

   - **`--port`**: Port where the game server is running (default: `8765`).
   - **`--bot`**: Choose the type of bot to use. Possible values:
     - `"RandomBot"`: Example bot implementation which plays random
     - `"MyBot"`: Template for your custom bot

## Create Your Own Bot

1. Create a copy of `mybot.go` in the `internal/bot` directory (e.g.
   `internal/bot/mycustombot.go`)
2. Rename all instances of `MyBot` in your new file to your bot's name
   (e.g. `MyCustomBot`)
3. Implement your bot logic in the `Run()` method
4. Add your bot in the `internal/bot/factory.go` file in the `NewBot()` method
5. Run the client with your bot's name using the `--bot MyCustomBot` flag

## Project Structure

```
│   go.mod
│   go.sum
│   README.md
├───cmd
│       main.go            // Main entry point for the application
└───internal
    ├───bot
    │       bot.go         // Base implementation of the bot
    │       factory.go     // Bot factory for creating bots
    │       mybot.go       // Template for creating a custom bot
    │       randombot.go   // Example bot implementation
    ├───model
    │       request.go     // Defines the requests for the gameserver
    │       response.go    // Defines the responses from the gameserver
    └───websocket
            client.go      // WebSocket client for communication
```
