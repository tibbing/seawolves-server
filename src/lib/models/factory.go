package models

import "fmt"

// Factory model
type Factory struct {
	modelImpl
	OwnerID                 string
	FactoryTypeID           string
	PortID                  string
	Storage                 Resource
	UpdatedAt               int
	ProductionSpeedModifier float32
	LocationID              int
}

// NewFactory Creates a new factory
func NewFactory(currentMap Map, factoryTypeID string, portID string, productionSpeedModifier float32, locationID int, ownerID string) *Factory {
	resourceTypeID := currentMap.FactoryTypes[factoryTypeID].ResourceTypeID
	resourceType := currentMap.ResourceTypes[resourceTypeID]

	result := &Factory{
		OwnerID:                 ownerID,
		FactoryTypeID:           factoryTypeID,
		PortID:                  portID,
		UpdatedAt:               0,
		Storage:                 *NewResource(0, resourceType),
		ProductionSpeedModifier: productionSpeedModifier,
		LocationID:              locationID,
	}
	result.SetRandomID()
	log.Infof("Creating factory %s", result.String())

	return result
}

// UpdateStorage Updates storage of factory based on number of days passed
func (x *Factory) UpdateStorage(currentMap Map, day int) {
	elapsedDays := day - x.UpdatedAt
	factoryType := currentMap.FactoryTypes[x.FactoryTypeID]
	produced := factoryType.ProductionSpeedModifier * x.ProductionSpeedModifier * float32(elapsedDays)
	x.UpdatedAt = day
	x.Storage.Amount += produced
	log.Infof("Adding %v to storage in factory %s, new value is %f %s. Elapsed days: %v",
		produced, x.String(), x.Storage.Amount, x.Storage.Type.GetID(), elapsedDays)
}

// GetID Gets the ID
func (x *Factory) GetID() string {
	return x.ID
}

func (x *Factory) String() string {
	return fmt.Sprintf("ID: %s, Type: %s PortID: %s, LocationID: %v, OwnerID: %s",
		x.GetID(), x.FactoryTypeID, x.PortID, x.LocationID, x.OwnerID)
}
