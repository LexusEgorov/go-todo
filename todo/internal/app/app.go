package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/LexusEgorov/todo/internal/config"
	"github.com/LexusEgorov/todo/internal/server"
)

const (
	opNew = "App.New"
)

type App struct {
	server *server.Server
	logger *slog.Logger
}

func New(logger *slog.Logger, config *config.Config) (*App, error) {
	server, err := server.New(logger, *config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	return &App{
		server: server,
		logger: logger,
	}, nil
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
