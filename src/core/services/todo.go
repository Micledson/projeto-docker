package services

import (
	"github.com/google/uuid"
	"projeto-docker/src/core/domain/errors"
	"projeto-docker/src/core/domain/todo"
	"projeto-docker/src/core/interfaces/adapters"
	"projeto-docker/src/core/interfaces/usecases"
)

type todoService struct {
	adapter adapters.TodoAdapter
}

func NewTodoService(repository adapters.TodoAdapter) usecases.TodoUseCase {
	return &todoService{repository}
}

func (s *todoService) List() ([]todo.Todo, errors.Error) {
	return s.adapter.List()
}

func (s *todoService) FindByID(id uuid.UUID) (todo.Todo, errors.Error) {
	return s.adapter.FetchByID(id)
}

func (s *todoService) Create(newTodo todo.Todo) (todo.Todo, errors.Error) {
	return s.adapter.Insert(newTodo)
}

func (s *todoService) Update(id uuid.UUID, newTodo todo.Todo) errors.Error {
	return s.adapter.Update(id, newTodo)
}

func (s *todoService) EnableToDo(id uuid.UUID) errors.Error {
	_todo, err := todo.NewBuilder().WithID(id).WithIsActive(true).Build()
	if err != nil {
		return err
	}
	return s.adapter.ChangeStatus(_todo)
}
func (s *todoService) DisableToDo(id uuid.UUID) errors.Error {
	_todo, err := todo.NewBuilder().WithID(id).WithIsActive(false).Build()
	if err != nil {
		return err
	}
	return s.adapter.ChangeStatus(_todo)
}
