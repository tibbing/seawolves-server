package models

// Factory model
type Factory struct {
	modelImpl
	OwnerID                 string
	FactoryTypeID           string
	PortID                  string
	Storage                 Resource
	UpdatedAt               int16
	ProductionSpeedModifier float32
}

// NewFactory Creates a new factory
func NewFactory(currentMap Map, factoryTypeID string, portID string, productionSpeedModifier float32, ownerID string) *Factory {
	resourceTypeID := currentMap.FactoryTypes[factoryTypeID].ResourceTypeID
	resourceType := currentMap.ResourceTypes[resourceTypeID]

	result := &Factory{
		OwnerID:                 ownerID,
		FactoryTypeID:           factoryTypeID,
		PortID:                  portID,
		UpdatedAt:               0,
		Storage:                 *NewResource(0, resourceType),
		ProductionSpeedModifier: productionSpeedModifier,
	}
	result.SetRandomID()
	return result
}

// UpdateStorage Updates storage of factory based on number of days passed
func (m *Factory) UpdateStorage(currentMap Map, day int16) {
	elapsedDays := day - m.UpdatedAt
	factoryType := currentMap.FactoryTypes[m.FactoryTypeID]
	produced := factoryType.ProductionSpeedModifier * m.ProductionSpeedModifier * float32(elapsedDays)
	m.UpdatedAt = day
	m.Storage.Amount += produced
}
