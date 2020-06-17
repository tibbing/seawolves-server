package models

// Model Interface
type Model interface {
	GetID() string
	SetID(id string)
}

type modelImpl struct {
	id string
}

func (m *modelImpl) SetID(id string) {
	m.id = id
}
