package models

// Port model
type Port struct {
	modelImpl
	PortTypeID string
	OwnerID    string
	Factories  []Factory
	Shipyard   *Shipyard
}

// NewPort Creates a new port
func NewPort(portTypeID, ownerID string) *Port {
	result := &Port{
		PortTypeID: portTypeID,
		OwnerID:    ownerID,
		Factories:  []Factory{},
	}
	result.SetID(portTypeID)
	return result
}

// GetID Gets the ID
func (x *Port) GetID() string {
	return x.ID
}

// AddFactory Adds a factory
func (x *Port) AddFactory(factory Factory) {
	x.Factories = append(x.Factories, factory)
}

// CreateShipyard Sets the shipyard
func (x *Port) CreateShipyard(shipyard Shipyard) {
	x.Shipyard = &shipyard
}
