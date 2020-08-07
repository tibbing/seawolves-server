package models

import "fmt"

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
func NewPlayer(name, id string, playerType PlayerType) *Player {
	result := &Player{
		Name:   name,
		Type:   playerType,
		Day:    0,
		Gold:   10000,
		Fleets: []Fleet{},
	}
	result.SetID(id)
	return result
}

// GetID Gets the ID
func (u *Player) GetID() string {
	return u.Name
}

// MakeTransaction makes a gold transaction
func (u *Player) MakeTransaction(amount int) error {
	if u.Gold+amount < 0 {
		return fmt.Errorf("Player %s cannot afford %v, has %v", u.GetID(), amount, u.Gold)
	}
	u.Gold += amount
	return nil
}

// AddFleet adds a new fleet
func (u *Player) AddFleet(fleet Fleet) {
	u.Fleets = append(u.Fleets, fleet)
}
