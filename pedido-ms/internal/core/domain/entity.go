package domain

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type (
	EntityId interface {
		GenerateId()
		ID() string
	}

	EntityIdImp struct {
		id string
	}
)

func (e *EntityIdImp) SetId(value string) {
	e.id = value
}

func (e *EntityIdImp) GenerateId() {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)

	e.id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}

func (e *EntityIdImp) ID() string {
	return e.id
}
