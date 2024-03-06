package internal

import (
	"log"
	"net/http"
)

// StartServer starts the game server on the specified port.
func StartServer(port string) {
	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
