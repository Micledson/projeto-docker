package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"projeto-docker/src/core/domain"
	"projeto-docker/src/core/domain/errors"
	"projeto-docker/src/core/domain/todo"
	"projeto-docker/src/core/interfaces/adapters"
	"projeto-docker/src/infra/repository"
	"projeto-docker/src/infra/repository/postgres/query"
	"strconv"
	"time"
)

type todoRepository struct{}

func NewTodoRepository() adapters.TodoAdapter {
	return &todoRepository{}
}

var todosList []todo.Todo

func (r *todoRepository) List() ([]todo.Todo, errors.Error) {
	rows, err := repository.Queryx(query.Todo().Select().All())
	if err != nil {
		return todosList, nil
		//return nil, err
	}
	defer rows.Close()

	var todos []todo.Todo
	for rows.Next() {
		var serializedTodo = map[string]interface{}{}
		rows.MapScan(serializedTodo)
		todo, err := newTodoFromMapRows(serializedTodo)
		if err != nil {
			return todosList, nil
			//return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *todoRepository) FetchByID(id uuid.UUID) (todo.Todo, errors.Error) {
	rows, err := repository.Queryx(query.Todo().Select().FindByID(), id)
	if err != nil {
		for _, _todo := range todosList {
			if *_todo.ID() == id {
				return _todo, nil
			}

		}
		//return nil, err
	}
	//defer rows.Close()

	//if !rows.Next() {
	return nil, errors.NewNotFound("Todo not found")
	//}

	var serializedTodo = map[string]interface{}{}
	rows.MapScan(serializedTodo)

	_todo, err := newTodoFromMapRows(serializedTodo)
	if err != nil {
		return nil, err
	}

	return _todo, nil
}

func (r *todoRepository) Insert(newTodo todo.Todo) (todo.Todo, errors.Error) {
	rows, err := repository.Queryx(query.Todo().Insert().Insert(), newTodo.Description(), newTodo.IsActive())
	if err != nil {
		now := time.Now()
		todoBuild, buildErr := todo.NewBuilder().
			WithID(uuid.New()).
			WithDescription(newTodo.Description()).
			WithIsActive(newTodo.IsActive()).
			WithCreatedAt(&now).
			WithUpdatedAt(&now).
			Build()
		if buildErr != nil {
			return nil, buildErr
		}
		todosList = append(todosList, todoBuild)
		return todoBuild, nil
	}
	defer rows.Close()

	var id uuid.UUID
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, errors.New(err)
		}
	}

	_todo, err := r.FetchByID(id)
	if err != nil {
		return nil, err
	}
	return _todo, nil
}

func (r *todoRepository) Update(id uuid.UUID, newTodo todo.Todo) errors.Error {
	_, err := repository.Queryx(query.Todo().Update().Update(), newTodo.Description(), newTodo.ID())
	if err != nil {
		for index, _todo := range todosList {
			fmt.Println(_todo.ID())
			fmt.Println(id)
			if *_todo.ID() == id {
				now := time.Now()
				todoBuild, buildErr := todo.NewBuilder().
					WithID(id).
					WithDescription(newTodo.Description()).
					WithIsActive(_todo.IsActive()).
					WithCreatedAt(_todo.CreatedAt()).
					WithUpdatedAt(&now).
					Build()
				if buildErr != nil {
					return buildErr
				}
				todosList = append(todosList[:index], append([]todo.Todo{todoBuild}, todosList[index+1:]...)...)
			}
		}
		//return err
	}

	return nil
}

func (r *todoRepository) ChangeStatus(newTodo todo.Todo) errors.Error {
	_, err := repository.Queryx(query.Todo().Update().ChangeStatus(), newTodo.IsActive(), newTodo.ID())
	if err != nil {
		for index, _todo := range todosList {
			fmt.Println(_todo.ID())
			fmt.Println(newTodo.ID())
			fmt.Println(newTodo.ID() == _todo.ID())

			if *_todo.ID() == *newTodo.ID() {
				now := time.Now()
				todoBuild, buildErr := todo.NewBuilder().
					WithID(*newTodo.ID()).
					WithDescription(_todo.Description()).
					WithIsActive(newTodo.IsActive()).
					WithCreatedAt(_todo.CreatedAt()).
					WithUpdatedAt(&now).
					Build()
				if buildErr != nil {
					return buildErr
				}
				todosList = append(todosList[:index], append([]todo.Todo{todoBuild}, todosList[index+1:]...)...)
			}
		}
		//return err
	}

	return nil
}

func newTodoFromMapRows(data map[string]interface{}) (todo.Todo, errors.Error) {
	var id uuid.UUID
	if parsedID, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		fmt.Println("erro massa:", err)
		return nil, errors.NewUnexpected()
	} else {
		id = parsedID
	}
	description := fmt.Sprint(data["description"])
	isActive, err := strconv.ParseBool(fmt.Sprint(data["is_active"]))
	if err != nil {
		return nil, errors.NewUnexpected()
	}

	createdAt := domain.ParseUTCTimestampToTime(fmt.Sprint(data["created_at"]))
	updatedAt := domain.ParseUTCTimestampToTime(fmt.Sprint(data["updated_at"]))

	_todo, err := todo.NewBuilder().
		WithID(id).
		WithDescription(description).
		WithIsActive(isActive).
		WithCreatedAt(createdAt).
		WithUpdatedAt(updatedAt).
		Build()
	if err != nil {
		return nil, errors.New(err)
	}
	return _todo, nil
}
