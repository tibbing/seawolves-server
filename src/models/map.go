package models

// Map model
type Map struct {
	modelImpl
	Name   string
	Width  int16
	Height int16
	Ports  []*Port
}

// NewMap Creates a new map
func NewMap(name string, width, height int16) *Map {
	result := &Map{
		Name:   name,
		Width:  width,
		Height: height,
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *Map) GetID() string {
	return u.Name
}

// AddPort Adds a port
func (u *Map) AddPort(port *Port) {
	u.Ports = append(u.Ports, port)
}
