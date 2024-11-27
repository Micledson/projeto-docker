package dicontainer

import (
	"projeto-docker/src/core/interfaces/usecases"
	"projeto-docker/src/core/services"
	"projeto-docker/src/infra/repository/postgres"
)

func TodosUseCase() usecases.TodoUseCase {
	repo := postgres.NewTodoRepository()
	return services.NewTodoService(repo)
}
