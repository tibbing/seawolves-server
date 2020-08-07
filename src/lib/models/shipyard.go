package models

import "lib/ext"

// Shipyard model
type Shipyard struct {
	Level               int
	ShipsTypesAvailable map[string]bool
}

// NewShipyard Creates a new Shipyard
func NewShipyard(currentMap Map, level int, shipsTypesAvailable []string) *Shipyard {
	shipTypesMap := make(map[string]bool)
	for k := range currentMap.ShipTypes {
		shipTypesMap[k] = ext.Contains(shipsTypesAvailable, k)
	}

	result := &Shipyard{
		Level:               level,
		ShipsTypesAvailable: shipTypesMap,
	}
	return result
}

// AddShip Adds an available ship
func (x *Shipyard) AddShip(shipTypeID string) {
	x.ShipsTypesAvailable[shipTypeID] = true
}

// RemoveShip Removes an available ship
func (x *Shipyard) RemoveShip(shipTypeID string) {
	x.ShipsTypesAvailable[shipTypeID] = false
}

// HasShip returns whether a ship type is available
func (x *Shipyard) HasShip(shipTypeID string) bool {
	return x.ShipsTypesAvailable[shipTypeID] == true
}
