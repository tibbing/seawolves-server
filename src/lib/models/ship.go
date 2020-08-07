package models

// Ship model
type Ship struct {
	modelImpl
	Name    string
	Type    string
	Storage []Resource
}

// NewShip Creates a new ship
func NewShip(name string, shipType string) *Ship {
	result := &Ship{
		Name: name,
		Type: shipType,
	}
	result.SetRandomID()
	return result
}

// GetID Gets the ID
func (u *Ship) GetID() string {
	return u.Name
}
