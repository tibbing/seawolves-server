package models

// Position model
type Position struct {
	X int8
	Y int8
}

// NewPosition Creates a new position
func NewPosition(x, y int8) *Position {
	result := &Position{
		X: x,
		Y: y,
	}
	return result
}
