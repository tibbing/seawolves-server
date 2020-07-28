package models

// Port model
type Port struct {
	PortTypeID string
	OwnerID    string
	Factories  []Factory
}

// NewPort Creates a new port
func NewPort(portTypeID, ownerID string) *Port {
	result := &Port{
		PortTypeID: portTypeID,
		OwnerID:    ownerID,
		Factories:  []Factory{},
	}
	return result
}

// AddFactory Adds a factory
func (x *Port) AddFactory(factory Factory) {
	x.Factories = append(x.Factories, factory)
}
