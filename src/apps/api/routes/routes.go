package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"projeto-docker/src/apps/api/docs"
	"projeto-docker/src/utils"
)

type Router interface {
	Load(*echo.Group)
}

type router struct {
}

func New() Router {
	return &router{}
}

// @title Telemetria API
// @version 1.0
// @description Telemetria API
// @contact.name OOPS TELECOM
// @contact.email micledson.dev@gmail.com
// @BasePath /api/telemedicina
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func (r *router) Load(rootEndpoint *echo.Group) {
	docs.SwaggerInfo.BasePath = utils.GetenvWithDefault("SWAGGER_BASE_PATH", "/api")

	r.LoadDocs(rootEndpoint)

	NewTodoRouter().Load(rootEndpoint)
}

func (r *router) LoadDocs(group *echo.Group) {
	group.GET("/docs/*", echoSwagger.WrapHandler)
	group.GET("/docs", func(c echo.Context) error {
		return c.Redirect(301, "/api/docs/index.html")
	})
}
