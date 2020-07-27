package models

// Ship model
type Ship struct {
	modelImpl
	ID      string
	Name    string
	OwnerID string
	Type    ShipType
	Storage []Resource
}

// NewShip Creates a new ship
func NewShip(id, name, ownerID string, shipType ShipType) *Ship {
	result := &Ship{
		ID:      id,
		Name:    name,
		OwnerID: ownerID,
		Type:    shipType,
	}
	result.SetID(id)
	return result
}

// GetID Gets the ID
func (u *Ship) GetID() string {
	return u.Name
}
