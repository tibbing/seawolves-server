package models

// Map model
type Map struct {
	modelImpl
	Name          string
	Width         int16
	Height        int16
	Ports         map[string]*Port
	ResourceTypes map[string]ResourceType
	FactoryTypes  map[string]FactoryType
}

// NewMap Creates a new map
func NewMap(name string, width, height int16) *Map {
	result := &Map{
		Name:          name,
		Width:         width,
		Height:        height,
		Ports:         map[string]*Port{},
		ResourceTypes: map[string]ResourceType{},
		FactoryTypes:  map[string]FactoryType{},
	}
	result.SetID(name)
	return result
}

// GetID Gets the ID
func (u *Map) GetID() string {
	return u.Name
}

// AddPort Adds a port
func (u *Map) AddPort(port *Port) {
	u.Ports[port.GetID()] = port
}

// AddResourceType Adds a resource type
func (u *Map) AddResourceType(resourceType ResourceType) {
	u.ResourceTypes[resourceType.GetID()] = resourceType
}

// AddFactoryType Adds a factory type
func (u *Map) AddFactoryType(factoryType FactoryType) {
	u.FactoryTypes[factoryType.GetID()] = factoryType
}

// GetPorts Lists all ports
func (u *Map) GetPorts() []string {
	keys := make([]string, 0, len(u.Ports))
	for k := range u.Ports {
		keys = append(keys, k)
	}
	return keys
}

// GetResourceTypes Lists all resource types
func (u *Map) GetResourceTypes() []string {
	keys := make([]string, 0, len(u.ResourceTypes))
	for k := range u.ResourceTypes {
		keys = append(keys, k)
	}
	return keys
}

// GetFactoryTypes Lists all factory types
func (u *Map) GetFactoryTypes() []string {
	keys := make([]string, 0, len(u.FactoryTypes))
	for k := range u.FactoryTypes {
		keys = append(keys, k)
	}
	return keys
}
