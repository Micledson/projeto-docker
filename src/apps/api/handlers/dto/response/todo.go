package response

import (
	"github.com/google/uuid"
	"projeto-docker/src/core/domain/todo"
	"time"
)

type Todo struct {
	ID          *uuid.UUID `json:"id"`
	Description string     `json:"description,omitempty"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type todoBuilder struct{}

func TodoBuilder() *todoBuilder {
	return &todoBuilder{}
}

func (*todoBuilder) BuildFromDomain(data todo.Todo) *Todo {
	return &Todo{
		data.ID(),
		data.Description(),
		data.IsActive(),
		data.CreatedAt(),
		data.UpdatedAt(),
	}
}
