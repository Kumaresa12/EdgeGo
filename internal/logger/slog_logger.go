package logger

import (
	"log/slog"
	"os"
)

type SLogger struct {
	logger *slog.Logger
}

func New() Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	return &SLogger{
		logger: slog.New(handler),
	}
}

func (l *SLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *SLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *SLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
