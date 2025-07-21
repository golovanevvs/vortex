package zlog

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func Init() {
	Logger = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()
}

func ParseLogLevel(level string) (zerolog.Level, error) {
	return zerolog.ParseLevel(level)
}
