package models

// ResourceType model
type ResourceType struct {
	modelImpl
	Name         string
	InitialValue int8
	Weight       int8
}

// NewResourceType Creates a new resource type
func NewResourceType(name string, initialValue int8, weight int8) ResourceType {
	result := ResourceType{
		Name:         name,
		InitialValue: initialValue,
		Weight:       weight,
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *ResourceType) GetID() string {
	return u.Name
}
