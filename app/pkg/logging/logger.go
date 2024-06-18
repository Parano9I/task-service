package logging

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func MustCreate(level string) *Logger {
	lvl := slog.Level(0)

	err := lvl.UnmarshalText([]byte(level))
	if err != nil {
		panic(err)
	}

	slogLogger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: lvl,
		}),
	)

	return &Logger{slogLogger}
}
