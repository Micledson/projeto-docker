package infra

import (
	"github.com/rs/zerolog"
	"projeto-docker/src/utils"
)

func Logger() zerolog.Logger {
	return utils.Logger().With().Str("layer", "infra").Logger()
}
