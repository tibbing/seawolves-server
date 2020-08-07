package models

// Player model
type Player struct {
	modelImpl
	Name   string
	Type   PlayerType
	Day    int
	Gold   int
	Fleets []Fleet
}

// NewPlayer Creates a new player
func NewPlayer(name string, playerType PlayerType) *Player {
	result := &Player{
		Name:   name,
		Type:   playerType,
		Day:    0,
		Gold:   10000,
		Fleets: []Fleet{},
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *Player) GetID() string {
	return u.Name
}
