package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/LexusEgorov/todo/internal/config"
	"github.com/LexusEgorov/todo/internal/server/handlers"
	"github.com/LexusEgorov/todo/internal/server/middleware"
)

const (
	opNew = "Server.New"
)

type Server struct {
	server *echo.Echo
	logger *slog.Logger
	config config.ServerConfig
}

func New(logger *slog.Logger, config config.Config) (*Server, error) {
	serverHandlers, err := handlers.New(logger, config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	middleware := middleware.New(logger)
	echoServer := echo.New()
	echoServer.Use(middleware.WithRecover, middleware.WithLogging)

	taskGroup := echoServer.Group("tasks")

	taskGroup.GET("/", serverHandlers.Task.GetAll)
	taskGroup.POST("/", serverHandlers.Task.Create)
	//check
	taskGroup.DELETE("/:id", serverHandlers.Task.Delete)
	taskGroup.GET("/:id", serverHandlers.Task.Get)
	taskGroup.POST("/update", serverHandlers.Task.Update)

	userGroup := echoServer.Group("users")
	userGroup.GET("/:id", serverHandlers.User.Get)
	userGroup.POST("/", serverHandlers.User.Update)

	//check
	userGroup.DELETE("/:id", serverHandlers.User.Delete)

	echoServer.POST("/register", serverHandlers.User.Register)
	echoServer.POST("/auth", serverHandlers.User.Auth)
	echoServer.POST("/refresh", serverHandlers.User.Refresh)

	return &Server{
		server: echoServer,
		logger: logger,
		config: config.Server,
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
