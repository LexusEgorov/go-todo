package handlers

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/LexusEgorov/todo/internal/config"
	"github.com/LexusEgorov/todo/internal/server/handlers/task"
	"github.com/LexusEgorov/todo/internal/server/handlers/user"
	"github.com/LexusEgorov/todo/internal/services"
)

const (
	opNew = "Handlers.New"
)

type UserHandler interface {
	Register(c echo.Context) error
	Auth(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type TaskHandler interface {
	Get(c echo.Context) error
	GetAll(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type Handlers struct {
	User UserHandler
	Task TaskHandler
}

func New(logger *slog.Logger, cfg config.DBConfig) (*Handlers, error) {
	services, err := services.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	return &Handlers{
		User: user.New(logger, services.User),
		Task: task.New(logger, services.Task),
	}, nil
}
