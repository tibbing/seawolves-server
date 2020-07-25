package models

// FactoryType model
type FactoryType struct {
	modelImpl
	Name           string
	ResourceTypeID string
}

// NewFactoryType Creates a new factory type
func NewFactoryType(name string, resourceTypeID string) FactoryType {
	result := FactoryType{
		Name:           name,
		ResourceTypeID: resourceTypeID,
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *FactoryType) GetID() string {
	return u.Name
}
