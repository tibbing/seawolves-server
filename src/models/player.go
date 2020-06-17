package models

// PlayerType Todo
type PlayerType int8

// Human player type
const (
	Human PlayerType = iota
	Ai
)

// Player model
type Player struct {
	modelImpl
	Name string
	Type PlayerType
}

// NewPlayer Creates a new player
func NewPlayer(name string, playerType PlayerType) *Player {
	result := &Player{
		Name: name,
		Type: playerType,
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *Player) GetID() string {
	return u.Name
}
