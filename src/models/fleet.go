package models

// Fleet model
type Fleet struct {
	modelImpl
	Name    string
	OwnerID string
	Ships   []Ship
}

// NewFleet Creates a new fleet
func NewFleet(name, ownerID string, ships []Ship) *Fleet {
	result := &Fleet{
		Name:    name,
		OwnerID: ownerID,
		Ships:   ships,
	}
	result.SetRandomID()
	return result
}

// GetID Gets the ID
func (u *Fleet) GetID() string {
	return u.id
}
