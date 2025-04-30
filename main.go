package main

import (
	"os"
)

func init() {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	setLogLevel(level)
}

func main() {
	logger := newLogger()

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")
	logger.Info("Info message with secret", "secret", "my-secret-value")
}
