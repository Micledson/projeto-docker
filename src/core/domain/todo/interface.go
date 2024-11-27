package todo

import (
	"github.com/google/uuid"
	"time"
)

type Todo interface {
	ID() *uuid.UUID
	Description() string
	IsActive() bool
	CreatedAt() *time.Time
	UpdatedAt() *time.Time
}
