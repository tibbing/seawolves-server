package models

// Port model
type Port struct {
	modelImpl
	Name      string
	OwnerID   string
	PositionX int16
	PositionY int16
}

// NewPort Creates a new port
func NewPort(name string, position Position) *Port {
	result := &Port{
		Name:      name,
		PositionX: position.X,
		PositionY: position.Y,
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *Port) GetID() string {
	return u.Name
}
