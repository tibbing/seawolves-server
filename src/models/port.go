package models

// Port model
type Port struct {
	PortTypeID string
	OwnerID    string
}

// NewPort Creates a new port
func NewPort(portTypeID, ownerID string) *Port {
	result := &Port{
		PortTypeID: portTypeID,
		OwnerID:    ownerID,
	}
	return result
}
