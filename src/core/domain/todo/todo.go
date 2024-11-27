package todo

import (
	"github.com/google/uuid"
	"time"
)

type todo struct {
	id          *uuid.UUID
	description string
	isActive    bool
	createdAt   *time.Time
	updatedAt   *time.Time
}

func (t todo) ID() *uuid.UUID {
	return t.id
}

func (t todo) Description() string {
	return t.description
}

func (t todo) IsActive() bool {
	return t.isActive
}

func (t todo) CreatedAt() *time.Time {
	return t.createdAt
}

func (t todo) UpdatedAt() *time.Time {
	return t.updatedAt
}
