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
	id string
}

func (m *modelImpl) SetID(id string) {
	m.id = id
}

func (m *modelImpl) SetRandomID() {
	m.id = xid.New()
}
