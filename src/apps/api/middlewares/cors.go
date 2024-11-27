package middlewares

import (
	"net/http"
	"projeto-docker/src/utils"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// verifyOrigin verifies if the request origin is included on the defined server
// allowed hosts.
func verifyOrigin(origin string) (bool, error) {
	allowedOrigins := strings.Split(utils.GetenvWithDefault("SERVER_ALLOWED_HOSTS", "*"), ",")
	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == "*" || origin == allowedOrigin {
			return true, nil
		}
	}
	return false, &echo.HTTPError{Code: 401, Message: "you're not allowed to access this API"}
}

// originInspectSkipper verifies the request context and skip the origin verification.
// It's useful to allow access for any origin when a route (e.g. public images routes)
// should be accessed by anyone.
func originInspectSkipper(context echo.Context) bool {
	return false
}

func CORSMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:         originInspectSkipper,
		AllowOriginFunc: verifyOrigin,
		AllowHeaders:    []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		},
	})
}
