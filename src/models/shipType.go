package models

// ShipType model
type ShipType struct {
	modelImpl
	Name             string
	Speed            int
	ResourceCapacity int
	ManCapacity      int
}

// NewShipType Creates a new ship type
func NewShipType(name string, speed, resourceCapacity, manCapacity int) ShipType {
	result := ShipType{
		Name:             name,
		Speed:            speed,
		ResourceCapacity: resourceCapacity,
		ManCapacity:      manCapacity,
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *ShipType) GetID() string {
	return u.Name
}
