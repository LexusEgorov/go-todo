package app

import (
	"context"
	"log/slog"

	"github.com/LexusEgorov/auth/internal/config"
	"github.com/LexusEgorov/auth/internal/server"
)

type App struct {
	server *server.Server
	logger *slog.Logger
}

func New(logger *slog.Logger, config *config.Config) *App {
	//TODO: init server
	return &App{
		logger: logger,
	}
}

func (a App) Run() {
	a.logger.Info("Starting app")
	go a.server.Run()
}

func (a App) Stop(ctx context.Context) {
	a.logger.Info("Stopping app...")

	doneCh := make(chan error)
	go func() {
		doneCh <- a.server.Stop(ctx)
	}()

	select {
	case err := <-doneCh:
		if err != nil {
			a.logger.Error(err.Error())
			return
		}

		a.logger.Info("App has been stopped gracefully")

	case <-ctx.Done():
		a.logger.Warn("App stopped forced")
	}
}
