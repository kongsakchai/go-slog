package goslog

import (
	"io"
	"log/slog"
)

func DefaultLogger(w io.Writer) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	return logger
}

func DefaultLoggerWithAttrs(w io.Writer, attrs []slog.Attr) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}).WithAttrs(attrs))

	return logger
}

func CustomLogger(w io.Writer) *slog.Logger {
	logger := slog.New(newTextHandler(w, &handlerOptions{
		level:      slog.LevelInfo,
		timeFormat: "[2006-01-02 15:04:05]",
	}))

	return logger
}

func CustomLoggerWithAttrs(w io.Writer, attrs []slog.Attr) *slog.Logger {
	logger := slog.New(newTextHandler(w, &handlerOptions{
		level:      slog.LevelInfo,
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
