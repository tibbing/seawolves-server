package models

// Factory model
type Factory struct {
	modelImpl
	OwnerID       string
	FactoryTypeID string
	PortID        string
}

// NewFactory Creates a new factory
func NewFactory(factoryTypeID string, portID string, ownerID string) *Factory {
	result := &Factory{
		OwnerID:       ownerID,
		FactoryTypeID: factoryTypeID,
		PortID:        portID,
	}
	result.SetRandomID()
	return result
}
