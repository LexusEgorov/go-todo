package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/LexusEgorov/auth/internal/config"
	"github.com/LexusEgorov/auth/internal/server/handlers"
	"github.com/labstack/echo/v4"
)

const (
	opNew = "Server.New"
)

type Server struct {
	server *echo.Echo
	logger *slog.Logger
	config *config.ServerConfig
}

func New(logger *slog.Logger, config *config.Config) (*Server, error) {
	handlers, err := handlers.New(logger, *config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}
	echoServer := echo.New()

	echoServer.POST("", handlers.User.Auth)
	echoServer.POST("", handlers.User.Auth)
	echoServer.POST("", handlers.User.Auth)
	echoServer.POST("", handlers.User.Auth)
	echoServer.POST("", handlers.User.Auth)

	return &Server{
		logger: logger,
		config: &config.Server,
	}, nil
}

func (s Server) Run() {
	serverAddr := fmt.Sprintf("%s:%d", s.config.Addr, s.config.Port)
	s.logger.Info(fmt.Sprintf("server is starting on %s", serverAddr))
	if err := s.server.Start(serverAddr); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error(err.Error())
		}
	}
}

func (s Server) Stop(ctx context.Context) error {
	s.logger.Info("stopping server...")
	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("Server.Stop: %w", err)
	}

	return nil
}
