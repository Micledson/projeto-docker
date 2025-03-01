package adapters

import (
	"github.com/google/uuid"
	"projeto-docker/src/core/domain/errors"
	"projeto-docker/src/core/domain/todo"
)

type TodoAdapter interface {
	List() ([]todo.Todo, errors.Error)
	FetchByID(uuid.UUID) (todo.Todo, errors.Error)
	Insert(todo.Todo) (todo.Todo, errors.Error)
	Update(uuid.UUID, todo.Todo) errors.Error
	ChangeStatus(todo.Todo) errors.Error
}
