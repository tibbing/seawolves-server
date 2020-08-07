package models

import (
	"github.com/rs/xid"
)

// Model Interface
type Model interface {
	GetID() string
	SetID(id string)
	SetRandomID()
}

type modelImpl struct {
	ID string
}

func (m *modelImpl) SetID(id string) {
	m.ID = id
}

func (m *modelImpl) SetRandomID() {
	m.ID = xid.New().String()
}
