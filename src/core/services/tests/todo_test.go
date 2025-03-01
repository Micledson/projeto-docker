package tests

import (
	mock_adapters "projeto-docker/src/core/services/tests/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"projeto-docker/src/core/domain/errors"
	"projeto-docker/src/core/domain/todo"
	core "projeto-docker/src/core/services"
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("Retorna todos os todos com sucesso", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)

		var expected []todo.Todo
		id, _ := uuid.NewUUID()
		now := time.Now()
		testTodo, _ := todo.NewBuilder().
			WithID(id).
			WithDescription("Descrição").
			WithIsActive(true).
			WithCreatedAt(&now).
			WithUpdatedAt(&now).
			Build()

		expected = append(expected, testTodo)
		IAdapter.EXPECT().List().Return(expected, nil)

		s := core.NewTodoService(IAdapter)
		todos, err := s.List()
		assert.Nil(t, err)
		assert.Equal(t, expected, todos)
	})

	t.Run("Retorna erro na listagem", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		expectedErr := errors.NewFromString("Error")
		IAdapter.EXPECT().List().Return(nil, expectedErr)

		s := core.NewTodoService(IAdapter)
		_, err := s.List()
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestFindByID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("Retorna todo por ID com sucesso", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id, _ := uuid.NewUUID()
		now := time.Now()
		expectedTodo, _ := todo.NewBuilder().
			WithID(id).
			WithDescription("Buscar Todo").
			WithIsActive(true).
			WithCreatedAt(&now).
			WithUpdatedAt(&now).
			Build()

		IAdapter.EXPECT().FetchByID(id).Return(expectedTodo, nil)

		s := core.NewTodoService(IAdapter)
		todoRes, err := s.FindByID(id)
		assert.Nil(t, err)
		assert.Equal(t, expectedTodo, todoRes)
	})

	t.Run("Retorna erro ao buscar todo por ID", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id, _ := uuid.NewUUID()
		expectedErr := errors.NewFromString("Erro ao buscar")
		IAdapter.EXPECT().FetchByID(id).Return(nil, expectedErr)

		s := core.NewTodoService(IAdapter)
		_, err := s.FindByID(id)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("Cria todo com sucesso", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id, _ := uuid.NewUUID()
		now := time.Now()
		newTodo, _ := todo.NewBuilder().
			WithID(id).
			WithDescription("Novo Todo").
			WithIsActive(true).
			WithCreatedAt(&now).
			WithUpdatedAt(&now).
			Build()

		IAdapter.EXPECT().Insert(newTodo).Return(newTodo, nil)

		s := core.NewTodoService(IAdapter)
		created, err := s.Create(newTodo)
		assert.Nil(t, err)
		assert.Equal(t, newTodo, created)
	})

	t.Run("Falha ao criar todo", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id, _ := uuid.NewUUID()
		now := time.Now()
		newTodo, _ := todo.NewBuilder().
			WithID(id).
			WithDescription("Novo Todo").
			WithIsActive(true).
			WithCreatedAt(&now).
			WithUpdatedAt(&now).
			Build()

		expectedErr := errors.NewFromString("Erro ao inserir")
		IAdapter.EXPECT().Insert(newTodo).Return(nil, expectedErr)

		s := core.NewTodoService(IAdapter)
		_, err := s.Create(newTodo)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("Atualiza todo com sucesso", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id, _ := uuid.NewUUID()
		now := time.Now()
		updateTodo, _ := todo.NewBuilder().
			WithID(id).
			WithDescription("Atualiza Todo").
			WithIsActive(true).
			WithCreatedAt(&now).
			WithUpdatedAt(&now).
			Build()

		IAdapter.EXPECT().Update(id, updateTodo).Return(nil)

		s := core.NewTodoService(IAdapter)
		err := s.Update(id, updateTodo)
		assert.Nil(t, err)
	})

	t.Run("Falha ao atualizar todo", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id, _ := uuid.NewUUID()
		now := time.Now()
		updateTodo, _ := todo.NewBuilder().
			WithID(id).
			WithDescription("Atualiza Todo").
			WithIsActive(true).
			WithCreatedAt(&now).
			WithUpdatedAt(&now).
			Build()

		expectedErr := errors.NewFromString("Erro ao atualizar")
		IAdapter.EXPECT().Update(id, updateTodo).Return(expectedErr)

		s := core.NewTodoService(IAdapter)
		err := s.Update(id, updateTodo)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestEnableToDo(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("Habilita todo com sucesso", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id := uuid.Must(uuid.NewRandom())

		IAdapter.EXPECT().ChangeStatus(gomock.Any()).Return(nil)

		s := core.NewTodoService(IAdapter)
		err := s.EnableToDo(id)
		assert.Nil(t, err)
	})

	t.Run("Falha ao habilitar todo", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id := uuid.Must(uuid.NewRandom())
		expectedErr := errors.NewFromString("Erro ao mudar status")
		IAdapter.EXPECT().ChangeStatus(gomock.Any()).Return(expectedErr)

		s := core.NewTodoService(IAdapter)
		err := s.EnableToDo(id)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestDisableToDo(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("Desabilita todo com sucesso", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id := uuid.Must(uuid.NewRandom())

		IAdapter.EXPECT().ChangeStatus(gomock.Any()).Return(nil)

		s := core.NewTodoService(IAdapter)
		err := s.DisableToDo(id)
		assert.Nil(t, err)
	})

	t.Run("Falha ao desabilitar todo", func(t *testing.T) {
		IAdapter := mock_adapters.NewMockTodoAdapter(controller)
		id := uuid.Must(uuid.NewRandom())
		expectedErr := errors.NewFromString("Erro ao mudar status")
		IAdapter.EXPECT().ChangeStatus(gomock.Any()).Return(expectedErr)

		s := core.NewTodoService(IAdapter)
		err := s.DisableToDo(id)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}
