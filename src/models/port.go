package models

// Port model
type Port struct {
	modelImpl
	Name     string
	OwnerID  string
	Position string
}

// NewPort Creates a new port
func NewPort(name, ownerID, position string) *Port {
	result := &Port{
		Name:     name,
		OwnerID:  ownerID,
		Position: position,
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *Port) GetID() string {
	return u.Name
}
