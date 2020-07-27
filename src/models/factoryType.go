package models

// FactoryType model
type FactoryType struct {
	modelImpl
	Name                    string
	ResourceTypeID          string
	ProductionSpeedModifier float32
}

// NewFactoryType Creates a new factory type
func NewFactoryType(name string, resourceTypeID string, productionSpeedModifier float32) FactoryType {
	result := FactoryType{
		Name:                    name,
		ResourceTypeID:          resourceTypeID,
		ProductionSpeedModifier: productionSpeedModifier,
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *FactoryType) GetID() string {
	return u.Name
}
