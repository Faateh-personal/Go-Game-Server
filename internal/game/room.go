package game

import "sync"

// Room represents a game room where players can join and play.
type Room struct {
	ID      string
	Players []*Player
	mutex   sync.Mutex
}

// NewRoom creates a new game room with a unique ID.
func NewRoom(id string) *Room {
	return &Room{
		ID: id,
	}
}

// AddPlayer adds a player to the room.
func (r *Room) AddPlayer(player *Player) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.Players = append(r.Players, player)
}
