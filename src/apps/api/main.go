package main

import (
	"fmt"
	"log"
	"projeto-docker/src/apps/api/middlewares"
	"projeto-docker/src/apps/api/routes"
	"projeto-docker/src/utils"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	godotenv.Load(".env")
	api := NewAPI(getServerHostAndPort())
	api.Serve()
}

func getServerHostAndPort() (string, int) {
	host := utils.GetenvWithDefault("SERVER_HOST", "0.0.0.0")
	portStr := utils.GetenvWithDefault("PORT", "8000")
	var port int
	if v, err := strconv.Atoi(portStr); err != nil {
		log.Fatal("The server port env variable must be a number (e.g 8000)")
	} else {
		port = v
	}
	return host, port
}

type API interface {
	Serve()
}

type api struct {
	host   string
	port   int
	server *echo.Echo
}

// @title Todo API
// @version 1.0
// @description Todo API
// @contact.name Projeto - docker
// @contact.email projeto@email.com
// @BasePath /api
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func NewAPI(host string, port int) API {
	server := echo.New()
	return &api{host, port, server}
}

func (a *api) Serve() {
	a.setupMiddlewares()
	a.loadRoutes()
	a.start()
}

func (a *api) setupMiddlewares() {
	a.server.Use(middleware.Recover())
	a.server.Use(middlewares.LoggerMiddleware())
	a.server.Use(middlewares.CORSMiddleware())
}

func (a *api) rootEndpoint() *echo.Group {
	return a.server.Group("/api")
}

func (a *api) loadRoutes() {
	manager := routes.New()
	manager.Load(a.rootEndpoint())
}

func (a *api) start() {
	address := fmt.Sprintf("%s:%d", a.host, a.port)
	if err := a.server.Start(address); err != nil {
		a.server.Logger.Fatal(err)
	}
}
