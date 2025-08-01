package app

import (
	"context"
	"log/slog"

	"github.com/LexusEgorov/todo/internal/config"
)

type App struct {
	//TODO: add server
	logger *slog.Logger
}

func New(logger *slog.Logger, config *config.Config) (*App, error) {
	//TODO: init server
	return &App{
		logger: logger,
	}, nil
}

func (a App) Run() {
	//TODO: start server
}

func (a App) Stop(ctx context.Context) {
	a.logger.Info("Stopping app...")

	doneCh := make(chan error)
	go func() {
		//TODO: stop server
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
