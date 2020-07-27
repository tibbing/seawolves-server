package models

import "fmt"

// Resource model
type Resource struct {
	Type   ResourceType
	Amount float32
}

// NewResource Creates a new resource
func NewResource(amount float32, resourceType ResourceType) *Resource {
	result := &Resource{
		Amount: amount,
		Type:   resourceType,
	}
	return result
}

func (t *Resource) String() string {
	return fmt.Sprintf("%f %s", t.Amount, t.Type.GetID())
}
