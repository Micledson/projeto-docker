package request

import (
	"github.com/google/uuid"
	"projeto-docker/src/core/domain/errors"
	"projeto-docker/src/core/domain/todo"
	"projeto-docker/src/utils/validator"
)

type Todo struct {
	ID          *uuid.UUID `json:"id"`
	Description string     `json:"description"`
	IsActive    bool       `json:"is_active"`
}

func (t *Todo) ToDomain() (todo.Todo, errors.Error) {
	builder := todo.NewBuilder()

	if validator.IsUUIDValid(*t.ID) {
		builder.WithID(*t.ID)
	}

	_todo, err := builder.
		WithDescription(t.Description).
		WithIsActive(t.IsActive).
		Build()
	if err != nil {
		return nil, err
	}

	return _todo, nil
}
