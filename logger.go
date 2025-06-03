package main

import (
	"io"
	"log/slog"
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

func defaultLogger(w io.Writer) *slog.Logger {
	logger = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: logLevel,
	}))

	return logger
}

func defaultLoggerWithAttrs(w io.Writer, attrs []slog.Attr) *slog.Logger {
	logger = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: logLevel,
	}).WithAttrs(attrs))

	return logger
}

func customLogger(w io.Writer) *slog.Logger {
	logger = slog.New(newTextHandler(w, &handlerOptions{
		level:      logLevel,
		timeFormat: "[2006-01-02 15:04:05]",
	}))

	return logger
}

func customLoggerWithAttrs(w io.Writer, attrs []slog.Attr) *slog.Logger {
	logger = slog.New(newTextHandler(w, &handlerOptions{
		level:      logLevel,
		timeFormat: "[2006-01-02 15:04:05]",
	}).WithAttrs(attrs))

	return logger
}

func replaceHandler(groups []string, a slog.Attr) slog.Attr {
	if a.Key == "secret" && a.Value.String() != "" {
		return slog.String("secret", "*********")
	}

	return a
}
