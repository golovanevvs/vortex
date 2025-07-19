package zlog

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func Init(level string) (err error) {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	Logger = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller().
		Logger()
	return
}
