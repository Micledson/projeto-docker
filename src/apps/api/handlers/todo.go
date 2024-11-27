package handlers

import (
	"net/http"
	"projeto-docker/src/apps/api/handlers/dto/request"
	"projeto-docker/src/apps/api/handlers/dto/response"
	"projeto-docker/src/core/interfaces/usecases"
)

type Todo interface {
	List(RichContext) error
	GetByID(RichContext) error
	Create(RichContext) error
	Update(RichContext) error
	RestoreToDo(RichContext) error
	DeleteToDo(RichContext) error
}

type todo struct {
	service usecases.TodoUseCase
}

func NewTodo(service usecases.TodoUseCase) Todo {
	return &todo{service}
}

// List
// @ID Todo.List
// @Summary Lista todos os ToDos ativos.
// @Description Esta rota retorna todas os ToDos.
// @Tags Todo
// @Produce json
// @Success 200 {array} response.Todo "Requisição realizada com sucesso."
// @Failure 422 {object} response.ErrorMessage "Algum dos dados informados não pôde ser processado. Verifique os dados fornecidos."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /todo [get]
func (t todo) List(context RichContext) error {
	todos, err := t.service.List()
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}

	var serializedTodos []response.Todo
	for _, todo := range todos {
		serializedTodos = append(serializedTodos, *response.TodoBuilder().BuildFromDomain(todo))
	}

	return context.JSON(http.StatusOK, serializedTodos)
}

// GetByID
// @ID Todo.GetByID
// @Summary Retorna os detalhes de um ToDo
// @Description Esta rota retorna todas os ToDos.
// @Tags Todo
// @Param id path string  true  "UUID de um ToDo"
// @Produce json
// @Success 200 {array} response.Todo "Requisição realizada com sucesso."
// @Failure 422 {object} response.ErrorMessage "Algum dos dados informados não pôde ser processado. Verifique os dados fornecidos."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /todo/{id} [get]
func (t todo) GetByID(context RichContext) error {
	id, valErr := GetIDFromPathParam("id", context)
	if valErr != nil {
		return valErr
	}

	_todo, err := t.service.FindByID(*id)
	if err != nil {
		response.ErrorBuilder().NewFromDomain(err)
	}

	return context.JSON(http.StatusOK, *response.TodoBuilder().BuildFromDomain(_todo))
}

// Create
// @ID Todo.Create
// @Summary Cria um ToDo.
// @Description Esta rota Cria um ToDo.
// @Tags Todo
// @Param json body request.Todo true "Parametros necessários para inserir um ToDo."
// @Produce json
// @Success 200 {array} response.Todo "Requisição realizada com sucesso."
// @Failure 422 {object} response.ErrorMessage "Algum dos dados informados não pôde ser processado. Verifique os dados fornecidos."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /todo [post]
func (t todo) Create(context RichContext) error {
	var requestDto request.Todo
	if bindErr := context.Bind(&requestDto); bindErr != nil {
		response.ErrorBuilder().NewUnsupportedMediaTypeError()
	}
	todoDomain, dtoErr := requestDto.ToDomain()
	if dtoErr != nil {
		response.ErrorBuilder().NewFromDomain(dtoErr)
	}
	_todo, err := t.service.Create(todoDomain)
	if err != nil {
		response.ErrorBuilder().NewFromDomain(err)
	}

	return context.JSON(http.StatusOK, response.TodoBuilder().BuildFromDomain(_todo))
}

// Update
// @ID Todo.Create
// @Summary Atualização do ToDo.
// @Description Esta rota Atualiza a descrição de um ToDo.
// @Tags Todo
// @Param id path string  true  "UUID de um ToDo"
// @Param json body request.Todo true "Parametros necessários para inserir um ToDo."
// @Success 204 "Requisição realizada com sucesso"
// @Failure 422 {object} response.ErrorMessage "Algum dos dados informados não pôde ser processado. Verifique os dados fornecidos."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /todo/{id} [put]
func (t todo) Update(context RichContext) error {
	id, valErr := GetIDFromPathParam("id", context)
	if valErr != nil {
		return valErr
	}

	var requestDto request.Todo
	if bindErr := context.Bind(&requestDto); bindErr != nil {
		response.ErrorBuilder().NewUnsupportedMediaTypeError()
	}

	requestDto.ID = *id
	todoDomain, dtoErr := requestDto.ToDomain()
	if dtoErr != nil {
		response.ErrorBuilder().NewFromDomain(dtoErr)
	}
	err := t.service.Update(todoDomain)
	if err != nil {
		response.ErrorBuilder().NewFromDomain(err)
	}

	return context.NoContent(http.StatusNoContent)
}

// RestoreToDo
// @ID Todo.RestoreToDo
// @Summary Restora um ToDo.
// @Description Esta rota Restora um ToDo deletado.
// @Tags Todo
// @Param id path string  true  "UUID de um ToDo"
// @Success 204 "Requisição realizada com sucesso"
// @Failure 422 {object} response.ErrorMessage "Algum dos dados informados não pôde ser processado. Verifique os dados fornecidos."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /todo/restore/{id} [put]
func (t todo) RestoreToDo(context RichContext) error {
	id, valErr := GetIDFromPathParam("id", context)
	if valErr != nil {
		return valErr
	}

	err := t.service.EnableToDo(*id)
	if err != nil {
		response.ErrorBuilder().NewFromDomain(err)
	}

	return context.NoContent(http.StatusNoContent)
}

// DeleteToDo
// @ID Todo.DeleteToDo
// @Summary Deleta um ToDo.
// @Description Esta rota Deleta um ToDo.
// @Tags Todo
// @Param id path string  true  "UUID de um ToDo"
// @Success 204 "Requisição realizada com sucesso"
// @Failure 422 {object} response.ErrorMessage "Algum dos dados informados não pôde ser processado. Verifique os dados fornecidos."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /todo/{id} [delete]
func (t todo) DeleteToDo(context RichContext) error {
	id, valErr := GetIDFromPathParam("id", context)
	if valErr != nil {
		return valErr
	}

	err := t.service.DisableToDo(*id)
	if err != nil {
		response.ErrorBuilder().NewFromDomain(err)
	}

	return context.NoContent(http.StatusNoContent)
}
