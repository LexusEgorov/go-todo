package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/LexusEgorov/auth/internal/config"
)

type Server struct {
	server *echo.Echo
	logger *slog.Logger
	config config.ServerConfig
}

func New(logger *slog.Logger, config config.ServerConfig) *Server {
	server := echo.New()

	/*Auth*/
	//Регистрация
	//Авторизация
	//Проверка токена

	return &Server{
		server: server,
		logger: logger,
		config: config,
	}
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
		return fmt.Errorf("Server.Stop: %v", err)
	}

	return nil
}
