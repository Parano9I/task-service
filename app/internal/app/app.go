package app

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"task-service/internal/config"
	"task-service/pkg/logging"
)

type App struct {
	logger     *logging.Logger
	cfg        *config.Config
	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(logger *logging.Logger, config *config.Config) (*App, error) {
	router := httprouter.New()

	return &App{
		logger: logger,
		cfg:    config,
		router: router,
	}, nil
}

func (a *App) Start() {
	address := fmt.Sprintf("%s:%s", a.cfg.Listen.Host, a.cfg.Listen.Port)

	a.logger.Info(fmt.Sprintf("Starting server: %s", address))

	a.httpServer = &http.Server{
		Addr:         address,
		Handler:      a.router,
		ReadTimeout:  a.cfg.Listen.RequestTimeout,
		WriteTimeout: a.cfg.Listen.RequestTimeout,
	}

	if err := a.httpServer.ListenAndServe(); err != nil {
		a.logger.Error("Failed to start server", "error", err)
	}

	a.logger.Error("Server stopped")
}
