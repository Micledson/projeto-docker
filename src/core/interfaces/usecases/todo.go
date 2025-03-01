package usecases

import (
	"github.com/google/uuid"
	"projeto-docker/src/core/domain/errors"
	"projeto-docker/src/core/domain/todo"
)

type TodoUseCase interface {
	List() ([]todo.Todo, errors.Error)
	FindByID(uuid.UUID) (todo.Todo, errors.Error)
	Create(todo.Todo) (todo.Todo, errors.Error)
	Update(uuid.UUID, todo.Todo) errors.Error
	EnableToDo(uuid.UUID) errors.Error
	DisableToDo(uuid.UUID) errors.Error
}
