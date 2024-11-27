package routes

import (
	"github.com/labstack/echo/v4"
	"projeto-docker/src/apps/api/dicontainer"
	"projeto-docker/src/apps/api/handlers"
	"projeto-docker/src/apps/api/middlewares"
)

type todo struct {
	handler handlers.Todo
}

func NewTodoRouter() Router {
	service := dicontainer.TodosUseCase()
	handler := handlers.NewTodo(service)
	return &todo{handler}
}

func (c todo) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/todo")
	router.GET("", middlewares.EnhanceContext(c.handler.List))
	router.GET("/:id", middlewares.EnhanceContext(c.handler.GetByID))
	router.POST("", middlewares.EnhanceContext(c.handler.Create))
	router.PUT("/:id", middlewares.EnhanceContext(c.handler.Update))
	router.PUT("/restore/:id", middlewares.EnhanceContext(c.handler.RestoreToDo))
	router.DELETE("/:id", middlewares.EnhanceContext(c.handler.DeleteToDo))

}
