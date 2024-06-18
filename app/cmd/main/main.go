package main

import (
	"log/slog"
	"task-service/internal/app"
	"task-service/internal/config"
	"task-service/pkg/logging"
)

func main() {
	slog.Info("Config initializing...")
	cfg := config.GetConfig()

	slog.Info("Logger initializing...")
	logger := logging.MustCreate(cfg.LogLevel)

	app, err := app.NewApp(logger, cfg)
	if err != nil {
		logger.Error("Cannot initializing server", err)
	}

	app.Start()
}
