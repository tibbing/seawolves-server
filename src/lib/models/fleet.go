package models

// Fleet model
type Fleet struct {
	modelImpl
	Name    string
	OwnerID string
	Ships   []Ship
	Course  []Position
}

// NewFleet Creates a new fleet
func NewFleet(name, ownerID string, ships []Ship) *Fleet {
	result := &Fleet{
		Name:    name,
		OwnerID: ownerID,
		Ships:   ships,
		Course:  []Position{},
	}
	result.SetRandomID()
	return result
}

// SetCourse Sets the course for the fleet
func (u *Fleet) SetCourse(course []Position) {
	u.Course = course
}

// GetID Gets the ID
func (u *Fleet) GetID() string {
	return u.ID
}
