package player

// Player represents a player in the game.
type Player struct {
	ID   string
	Name string
	// Add more player properties here
}

// NewPlayer creates a new player with a unique ID and name.
func NewPlayer(id, name string) *Player {
	return &Player{
		ID:   id,
		Name: name,
	}
}
