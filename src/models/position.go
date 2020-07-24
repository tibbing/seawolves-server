package models

// Position model
type Position struct {
	X int16
	Y int16
}

// NewPosition Creates a new position
func NewPosition(x, y int16) Position {
	result := Position{
		X: x,
		Y: y,
	}
	return result
}
