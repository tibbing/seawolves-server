package models

import (
	"fmt"
	"time"
)

// Game model
type Game struct {
	modelImpl
	MapID   string
	Started int64
	Players []Player
	Ports   []Port
	Turn    string
}

// NewGame Creates a new game
func NewGame(mapID string, players []Player) *Game {
	result := &Game{
		MapID:   mapID,
		Started: time.Now().UTC().Unix(),
		Players: players,
		Ports:   []Port{},
		Turn:    players[0].GetID(),
	}
	result.SetRandomID()
	return result
}

// GetID Gets the ID
func (x *Game) GetID() string {
	return x.id
}

func (x *Game) String() string {
	unixTimeUTC := time.Unix(x.Started, 0) //gives unix time stamp in utc

	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format

	return fmt.Sprintf("%s %v", x.MapID, unitTimeInRFC3339)
}
