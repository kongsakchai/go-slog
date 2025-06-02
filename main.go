package main

import (
	"fmt"
	"log/slog"
	"os"
)

func init() {
	setLogLevel(os.Getenv("LOG_LEVEL"))
}

func main() {

	fmt.Println("Default Logger:")
	logTest(defaultLogger())

	fmt.Println("\nCustom Logger:")
	logTest(customLogger())
}

func logTest(l *slog.Logger) {
	l.Debug("Debug message", slog.String("key", "value"))
	l.Info("Info message", slog.String("key", "value"))
	l.Warn("Warning message", slog.String("key", "value"))
	l.Error("Error message", slog.String("key", "value"))

	w := l.With(slog.String("key", "value"))
	w.Info("Info message and with")
	w.Info("Info message and with attr", slog.String("key", "value2"))

	g := l.WithGroup("group1")
	g.Info("Info message in group", slog.String("key", "value3"))
}
