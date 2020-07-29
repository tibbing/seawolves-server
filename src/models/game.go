package models

import (
	"errors"
	"fmt"
	"time"
)

// Game model
type Game struct {
	modelImpl
	MapID   string
	Started int64
	Players map[string]*Player
	Ports   map[string]*Port
	Turn    string
}

// NewGame Creates a new game
func NewGame(mapID string, players []Player) *Game {

	result := &Game{
		MapID:   mapID,
		Started: time.Now().UTC().Unix(),
		Players: map[string]*Player{},
		Ports:   map[string]*Port{},
		Turn:    players[0].GetID(),
	}
	result.SetRandomID()
	log.Infof("Creating game %s", result.String())

	playersMap := make(map[string]*Player)
	for _, player := range players {
		log.Infof("Adding player %s to game %s", player.GetID(), result.GetID())
		playersMap[player.GetID()] = &player
	}
	result.Players = playersMap

	return result
}

// GetID Gets the ID
func (x *Game) GetID() string {
	return x.id
}

// AddPort Adds a port to the game
func (x *Game) AddPort(port *Port) error {
	x.Ports[port.PortTypeID] = port
	return nil
}

// UpdateForPlayer Updates the game for a given player and day
func (x *Game) UpdateForPlayer(currentMap *Map, playerID string, day int16) error {
	if day <= x.Players[playerID].Day {
		err := errors.New("Invalid day provided: " + string(day))
		return err
	}
	for _, port := range x.Ports {
		for _, factory := range port.Factories {
			if factory.OwnerID == playerID {
				factory.UpdateStorage(*currentMap, day)
			}
		}
	}
	return nil
}

func (x *Game) String() string {
	unixTimeUTC := time.Unix(x.Started, 0)                //gives unix time stamp in utc
	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format

	return fmt.Sprintf("ID: %s, MapID: %s Started: %s, Players: %v, Ports: %v, Turn: %s",
		x.GetID(), x.MapID, unitTimeInRFC3339, len(x.Players), len(x.Ports), x.Turn)
}
