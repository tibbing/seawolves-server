package models

// ResourceType model
type ResourceType struct {
	Name         string
	InitialValue int8
	Weight       int8
}

// NewResourceType Creates a new factory
func NewResourceType(name string, initialValue int8, weight int8) *ResourceType {
	result := &ResourceType{
		Name:         name,
		InitialValue: initialValue,
		Weight:       weight,
	}
	return result
}

// GetID Gets the ID
func (u *ResourceType) GetID() string {
	return u.Name
}
