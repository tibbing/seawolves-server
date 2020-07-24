package models

// FactoryType model
type FactoryType struct {
	Name           string
	ResourceTypeID string
}

// NewFactoryType Creates a new factory
func NewFactoryType(name string, resourceTypeID string) *FactoryType {
	result := &FactoryType{
		Name:           name,
		ResourceTypeID: resourceTypeID,
	}
	return result
}

// GetID Gets the ID
func (u *FactoryType) GetID() string {
	return u.Name
}
