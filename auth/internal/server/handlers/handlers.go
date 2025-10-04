package handlers

import (
	"fmt"
	"log/slog"

	"github.com/LexusEgorov/auth/internal/config"
	"github.com/LexusEgorov/auth/internal/server/handlers/token"
	"github.com/LexusEgorov/auth/internal/server/handlers/user"
	"github.com/LexusEgorov/auth/internal/services"
	"github.com/labstack/echo/v4"
)

const (
	opNew = "Handlers.New"
)

type UserHandler interface {
	Register(c echo.Context) error
	Auth(c echo.Context) error
}

type TokenHandler interface {
	Access(c echo.Context) error
	Block(c echo.Context) error
	Refresh(c echo.Context) error
}

type Handlers struct {
	User  UserHandler
	Token TokenHandler
}

func New(logger *slog.Logger, cfg config.Config) (*Handlers, error) {
	services, err := services.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	return &Handlers{
		User:  user.New(logger, services.User),
		Token: token.New(logger, services.Token),
	}, nil
}
