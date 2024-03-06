package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"game-server/internal/websocket"
)

func main() {
	// Set up a router
	r := mux.NewRouter()

	// Define a simple home handler
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Game Server!")
	})

	// Set up the WebSocket endpoint
	r.HandleFunc("/ws", websocket.HandleConnections)

	// Start the server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
