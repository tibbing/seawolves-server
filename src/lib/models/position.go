package models

// Position model
type Position struct {
	X int
	Y int
}

// NewPosition Creates a new position
func NewPosition(x, y int) Position {
	result := Position{
		X: x,
		Y: y,
	}
	return result
}
