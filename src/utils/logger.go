package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func Logger() zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatFieldValue = func(i interface{}) string {
		var log string
		if v, ok := i.(string); !ok {
			log = fmt.Sprintf("%#v", i)
		} else {
			log = v
		}
		if len(log) > 0 && log[0] == '"' {
			log = log[1:]
		}
		if len(log) > 0 && log[len(log)-1] == '"' {
			log = log[:len(log)-1]
		}
		return fmt.Sprintf("\"%s\"", log)
	}
	return zerolog.New(output).With().Timestamp().Logger()
}
