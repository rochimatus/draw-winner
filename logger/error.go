package logger

import "log/slog"

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Error(err error, msg string, args ...any) {
	args = append([]any{
		slog.String("error", err.Error()),
	}, args...)
	slog.Error(msg, args...)
}
