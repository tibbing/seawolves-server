package models

import (
	"errors"
	"fmt"
	"strconv"
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
		log.Infof("Adding player %s with ID %s to game %s", player.Name, player.GetID(), result.GetID())
		playersMap[player.GetID()] = &player
	}
	result.Players = playersMap

	return result
}

// GetID Gets the ID
func (x *Game) GetID() string {
	return x.ID
}

// AddPort Adds a port to the game
func (x *Game) AddPort(port *Port) error {
	x.Ports[port.PortTypeID] = port
	return nil
}

// UpdateForPlayer Updates the game for a given player and day
func (x *Game) UpdateForPlayer(currentMap *Map, playerID string, day int) error {
	log.Debugf("Updating game %s at day %v (current %v) for player %s", x.GetID(), day, x.Players[playerID].Day, playerID)
	if day <= x.Players[playerID].Day {
		err := errors.New("Invalid day provided: " + strconv.Itoa(day))
		return err
	}
	for i, port := range x.Ports {
		for j, factory := range port.Factories {
			if factory.OwnerID == playerID {
				factory.UpdateStorage(*currentMap, day)
				port.Factories[j] = factory
				x.Ports[i] = port
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
