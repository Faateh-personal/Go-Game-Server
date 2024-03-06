# Go-Based Game Server

## Overview

This project is a game server built in Go. It supports multiple game rooms, player matchmaking, and real-time communication using WebSockets. The server is designed to be scalable and efficient, making it suitable for a variety of multiplayer games.

## Features

- **Multiple Game Rooms**: Support for creating and managing multiple game rooms.
- **Player Matchmaking**: Algorithm for matching players based on skill level or other criteria.
- **Real-time Communication**: Real-time communication between clients and the server using WebSockets.
- **Game State Management**: Synchronization of game state across all clients in a room.
- **Authentication**: Secure authentication mechanism for player identification.
- **Logging**: Comprehensive logging for monitoring and debugging.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/username/game-server.git
   cd game-server
   ```

2. Install the dependencies:
   ```bash
   go mod tidy
   ```

### Running the Server

1. Start the server:
   ```bash
   go run cmd/server/main.go
   ```

2. The server will be running on `http://localhost:8080`.

## Usage

- Connect to the WebSocket endpoint at `ws://localhost:8080/ws` to start sending and receiving messages.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or create an issue for any bugs or feature requests.

