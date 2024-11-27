package middlewares

import (
	"github.com/labstack/echo/v4"
	"projeto-docker/src/apps/api/handlers"
)

func EnhanceContext(next handlers.RichHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		enhancedCtx, err := handlers.NewRichContext(ctx)
		if err != nil {
			return err
		}
		return next(enhancedCtx)
	}
}
