package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogError:    true,
		LogRemoteIP: true,
		LogURIPath:  true,
		LogURI:      true,
		LogStatus:   true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			var event *zerolog.Event
			event.Str("PATH", v.URIPath).Str("REMOTEIP", v.RemoteIP).Int("STATUS", v.Status)
			if v.Error != nil {
				event.Str("ERROR", v.Error.Error())
			}
			method := v.Method
			if v.Method == "OPTIONS" {
				method = "OPTS"
			}
			event.Msg(fmt.Sprintf("[%-5s]", method))
			return nil
		},
	})
}
