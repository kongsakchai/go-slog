package main

import (
	"log/slog"
	"os"
)

var logLevel slog.Level
var logger *slog.Logger

func setLogLevel(level string) {
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
}

func newLogger() *slog.Logger {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       logLevel,
		ReplaceAttr: replaceHandler,
	}))
	slog.SetDefault(logger)

	return logger
}

func replaceHandler(groups []string, a slog.Attr) slog.Attr {
	if a.Key == "secret" && a.Value.String() != "" {
		return slog.String("secret", "*********")
	}

	return a
}
