package models

import "fmt"

// Map model
type Map struct {
	modelImpl
	Name          string
	Width         int16
	Height        int16
	PortTypes     map[string]*PortType
	ResourceTypes map[string]ResourceType
	FactoryTypes  map[string]FactoryType
}

// NewMap Creates a new map
func NewMap(name string, width, height int16) *Map {
	result := &Map{
		Name:          name,
		Width:         width,
		Height:        height,
		PortTypes:     map[string]*PortType{},
		ResourceTypes: map[string]ResourceType{},
		FactoryTypes:  map[string]FactoryType{},
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (x *Map) GetID() string {
	return x.Name
}

// AddPortType Adds a port
func (x *Map) AddPortType(portType *PortType) {
	x.PortTypes[portType.GetID()] = portType
}

// AddResourceType Adds a resource type
func (x *Map) AddResourceType(resourceType ResourceType) {
	x.ResourceTypes[resourceType.GetID()] = resourceType
}

// AddFactoryType Adds a factory type
func (x *Map) AddFactoryType(factoryType FactoryType) {
	x.FactoryTypes[factoryType.GetID()] = factoryType
}

// GetPorts Lists all ports
func (x *Map) GetPorts() []string {
	keys := make([]string, 0, len(x.PortTypes))
	for k := range x.PortTypes {
		keys = append(keys, k)
	}
	return keys
}

// GetResourceTypes Lists all resource types
func (x *Map) GetResourceTypes() []string {
	keys := make([]string, 0, len(x.ResourceTypes))
	for k := range x.ResourceTypes {
		keys = append(keys, k)
	}
	return keys
}

// GetFactoryTypes Lists all factory types
func (x *Map) GetFactoryTypes() []string {
	keys := make([]string, 0, len(x.FactoryTypes))
	for k := range x.FactoryTypes {
		keys = append(keys, k)
	}
	return keys
}

func (x *Map) String() string {
	return fmt.Sprintf("Name: %s,\n Width: %v,\n Height: %v,\n PortTypes: %v,\n ResourceTypes: %v,\n FactoryTypes: %v\n",
		x.Name, x.Width, x.Height, len(x.PortTypes), len(x.ResourceTypes), len(x.FactoryTypes))
}
