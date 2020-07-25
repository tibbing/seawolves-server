package models

// ShipType model
type ShipType struct {
	modelImpl
	Name             string
	Speed            int8
	ResourceCapacity int8
	ManCapacity      int8
}

// NewShipType Creates a new ship type
func NewShipType(name string, speed, resourceCapacity, manCapacity int8) ShipType {
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
