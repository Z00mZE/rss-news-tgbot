package logger

import (
	"log/slog"
	"os"
	"strings"
)

type LogLevel string

const (
	LoggModeProd LogLevel = `prod`
	LoggModeTest LogLevel = `test`
	LoggModeDev  LogLevel = `dev`
)

var presets = map[LogLevel]slog.Handler{
	LoggModeProd: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
	LoggModeTest: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	LoggModeDev: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug},
	),
}

func NewSlogLogger(level LogLevel) *slog.Logger {
	loggerHandler, isExist := presets[level]

	if !isExist {
		loggerHandler = presets[LoggModeTest]
	}

	return slog.New(loggerHandler)
}

func ParseLogLevel(rawLogLevel string) LogLevel {
	parsedLvl := LogLevel(strings.ToLower(rawLogLevel))
	_, isExists := presets[parsedLvl]
	if !isExists {
		return parsedLvl
	}
	return parsedLvl
}
