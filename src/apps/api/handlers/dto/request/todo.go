package request

import (
	"projeto-docker/src/core/domain/errors"
	"projeto-docker/src/core/domain/todo"
)

type Todo struct {
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (t *Todo) ToDomain() (todo.Todo, errors.Error) {
	_todo, err := todo.NewBuilder().
		WithDescription(t.Description).
		WithIsActive(t.IsActive).
		Build()
	if err != nil {
		return nil, err
	}

	return _todo, nil
}

type UpdateTodo struct {
	Description string `json:"description"`
}

func (t *UpdateTodo) ToDomain() (todo.Todo, errors.Error) {
	_todo, err := todo.NewBuilder().
		WithDescription(t.Description).
		Build()
	if err != nil {
		return nil, err
	}

	return _todo, nil
}
