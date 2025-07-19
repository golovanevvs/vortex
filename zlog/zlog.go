package zlog

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func Init() {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		FormatLevel: func(i interface{}) string {
			return colorizeLevel(i.(string))
		},
	}
	Logger = zerolog.New(consoleWriter).
		With().
		Timestamp().
		Caller().
		Logger()
}

func ParseLogLevel(level string) (zerolog.Level, error) {
	return zerolog.ParseLevel(level)
}

func colorizeLevel(level string) string {
	switch level {
	case "debug":
		return "\x1b[36m" + level + "\x1b[0m" // Голубой
	case "info":
		return "\x1b[32m" + level + "\x1b[0m" // Зелёный
	case "warn":
		return "\x1b[33m" + level + "\x1b[0m" // Жёлтый
	case "error", "fatal", "panic":
		return "\x1b[31m" + level + "\x1b[0m" // Красный
	default:
		return level
	}
}
