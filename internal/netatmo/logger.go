package netatmo

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

// InitLogger initializes the package logger with the specified verbosity level
func InitLogger(verbose bool) {
	level := slog.LevelWarn
	if verbose {
		level = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	logger = slog.New(slog.NewTextHandler(os.Stderr, opts))
}

// GetLogger returns the package logger, initializing with defaults if needed
func GetLogger() *slog.Logger {
	if logger == nil {
		InitLogger(false)
	}
	return logger
}

// Debug logs a debug message
func Debug(msg string, args ...any) {
	GetLogger().Debug(msg, args...)
}

// Info logs an info message
func Info(msg string, args ...any) {
	GetLogger().Info(msg, args...)
}

// Warn logs a warning message
func Warn(msg string, args ...any) {
	GetLogger().Warn(msg, args...)
}

// Error logs an error message
func Error(msg string, args ...any) {
	GetLogger().Error(msg, args...)
}
