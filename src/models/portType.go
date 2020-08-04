package models

// PortType model
type PortType struct {
	modelImpl
	Name      string
	PositionX int
	PositionY int
}

// NewPortType Creates a new port type
func NewPortType(name string, position Position) *PortType {
	result := &PortType{
		Name:      name,
		PositionX: position.X,
		PositionY: position.Y,
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *PortType) GetID() string {
	return u.Name
}
