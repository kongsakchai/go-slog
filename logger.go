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

func defaultLogger() *slog.Logger {
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       logLevel,
		ReplaceAttr: replaceHandler,
	}))

	return logger
}

func customLogger() *slog.Logger {
	logger = slog.New(newTextHandler(os.Stdout, handlerOptions{
		level:       logLevel,
		replaceAttr: replaceHandler,
		timeFormat:  "[2006-01-02 15:04:05]",
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
