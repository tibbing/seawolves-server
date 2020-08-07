package models

// ResourceType model
type ResourceType struct {
	modelImpl
	Name         string
	InitialValue int
	Weight       int
}

// NewResourceType Creates a new resource type
func NewResourceType(name string, initialValue int, weight int) ResourceType {
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
