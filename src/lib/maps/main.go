package maps

import (
	"errors"
	"lib/models"
)

// GetMapByID returns a map instance by ID
func GetMapByID(id string) (models.Map, error) {
	switch id {
	case "Scandinavia":
		return *Scandinavia(), nil
	default:
		return models.Map{}, errors.New("Invalid map ID: " + id)
	}
}
